include "all_equal.mzn";

%%%%%%%%%%%%%
% VARAIBLES %
%%%%%%%%%%%%%

% days
int: nb_weeks;
int: nb_days = 7 * nb_weeks;
set of int: D = 1..nb_days;


% shift types
enum ShiftType = {
  JANW,
  JAEV,
  JAWH,
  SAEW,
  SAWH,
  TSPT,
  CALL,
  FREE
};

set of ShiftType: S;     
set of ShiftType: S_all = {JANW, JAEV, JAWH, SAEW, SAWH, TSPT, CALL};
set of ShiftType: S_full = S union {FREE};   
set of ShiftType: S_week = S intersect {JANW, SAEW, CALL};    
set of ShiftType: S_weekend = S intersect {JAWH, SAWH, TSPT};
set of ShiftType: S_weekday = S intersect {JAEV}; 


% personnel
enum PersonnelType = {
  JA,
  JA_F,
  SA,
  SA_F,
  SA_NEO,
  SA_F_NEO
};

int: nb_personnel;
set of 1..nb_personnel: P = 1..nb_personnel;   % employees
array[1..nb_personnel] of PersonnelType: T;    % mapping from employee to personnel type
array[1..nb_personnel] of set of D: F;         % granted holidays per employee 

% coverage requirements 
array[1..7, S_all] of int: R = [|1, 1, 0, 1, 0, 0, 1|      % Fri
                                 1, 0, 2, 1, 2, 1, 1|      % Sat
                                 1, 0, 2, 1, 2, 1, 1|      % Sun
                                 1, 1, 0, 1, 0, 0, 1|      % Mon
                                 1, 1, 0, 1, 0, 0, 1|      % Tue
                                 1, 1, 0, 1, 0, 0, 1|      % Wed
                                 1, 1, 0, 1, 0, 0, 1|];    % Thu 
                                 
% qualifications                                                                                                                                       
array[S_all] of set of PersonnelType: Q = [{JA},                            % JANW
                                           {JA},                            % JAEV
                                           {JA, JA_F},                      % JAWH
                                           {SA, SA_NEO},                    % SAEW
                                           {SA, SA_F, SA_NEO, SA_F_NEO},    % SAWH
                                           {SA_NEO, SA_F_NEO},              % TSPT
                                           {SA, SA_F, SA_NEO, SA_F_NEO}];   % CALL

 
                                                                                               
% model variables
array[S_all] of float: fairness_weight;
set of ShiftType: S_balance;
int: min_balance;

                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 
% decision variables X[p, s, d] == 1 -> employee p is assigned shift of type s on day d
array[1..nb_personnel, S_full, D] of var 0..1: X;

%%%%%%%%%%%%%%%
% CONSTRAINTS %
%%%%%%%%%%%%%%%


% HC1 - shift requirements met for every shift on every day
constraint forall(s in S, d in D)((sum(p in P)(X[p, s, d])) = R[1 + ((d-1) mod 7), s]);

% HC2 - only qualified personnel allowed
constraint forall(s in S, p in P, d in D)(X[p, s, d] = 1 -> T[p] in Q[s]);

% HC3 - exaclty one assignment per personnel member on every day
constraint forall(p in P, d in D)(sum(s in S_full)(X[p, s, d]) = 1);

% HC4 - same person for week shifts
constraint forall(w in 1..nb_weeks, s in S_week, p in P)(all_equal(X[p, s, 7*w-6..7*w]));


% HC5 - same person for weekend shifts
constraint forall(w in 1..nb_weeks, s in S_weekend, p in P)(X[p, s, 7*w-5] = X[p, s, 7*w-4]);


% HC6 - no assignments on free days
constraint forall(s in S, p in P, d in D)(X[p, s, d] = 1 -> not (d in F[p]));


% HC7 - no consecutive week assignements
constraint forall(p in P, s in S_week, w in 2..nb_weeks)(
             exists(s_prev in S)(X[p, s_prev, 7*w - 7] = 1) -> X[p, s, 7*w - 6] = 0
           );
           
                           
% HC8 - no consecutive weekday assignements
constraint forall(p in P, s in S_weekday, d in 2..nb_days)(
             exists(s_prev in S)(X[p, s_prev, d - 1] = 1) -> X[p, s, d] = 0
           );          


% optimization objective: fairness
array[S] of var 0..1.0: fairness_per_shift_type;
constraint forall(s in S)(fairness_per_shift_type[s] == max(p in P where T[p] in Q[s])(workload[s, p]) - min(p in P where T[p] in Q[s])(workload[s, p]));

array[S, P] of var 0..1.0: workload;
constraint forall(s in S, p in P)(workload[s, p] == (sum(d in D)(X[p, s, d])) / (card(D) - card(F[p])));

array[S] of var float: weighted_fairness;
constraint forall(s in S)(weighted_fairness[s] = fairness_per_shift_type[s] * fairness_weight[s]);

var 0..1.0: fairness_score;
constraint fairness_score == max(fairness_per_shift_type);


% optimization objective: balance

array[P, D] of var 0..nb_days: a;
constraint forall(p in P, d in D)(
             if d = 1
             then 
               if sum(s in S_balance)(X[p, s, d]) = 1
               then a[p, d] = 0
               else a[p, d] = 8
               endif
             else 
               if sum(s in S_balance)(X[p, s, d]) = 0 
               then a[p, d] = 1 + a[p, d-1]
               else a[p, d] = 0
               endif
             endif);
           
array[P, D] of var 0..nb_days: z;
constraint forall(p in P, d in D)(
             if d = nb_days
             then  z[p, d] = nb_days
             else 
               if sum(s in S_balance)(X[p, s, d]) = 0 /\ sum(s in S_balance)(X[p, s, d+1]) = 1
               then z[p, d] = a[p, d]
               else z[p, d] = nb_days
               endif
             endif
           );
           
constraint min(z) >= min_balance;

solve minimize fairness_score;

output 
       ["nb_days:"] ++ [show(nb_days)] ++ ["\n"] ++
       ["shift_types:"] ++ [show(s) ++ " " | s in S] ++ ["\n"] ++
       [
         if d == 1 then "assistant:" ++ show(p) ++ " " ++ "type:" ++ show(T[p]) ++ " " else "" endif ++
         if fix(X[p, FREE, d]) == 1 then "FREE " else "" endif ++
         if fix(X[p, JANW, d]) == 1 then "JANW " else "" endif ++
         if fix(X[p, SAEW, d]) == 1 then "SAEW " else "" endif ++
         if fix(X[p, JAEV, d]) == 1 then "JAEV " else "" endif ++
         if fix(X[p, JAWH, d]) == 1 then "JAWH " else "" endif ++
         if fix(X[p, SAWH, d]) == 1 then "SAWH " else "" endif ++
         if fix(X[p, TSPT, d]) == 1 then "TSPT " else "" endif ++
         if fix(X[p, CALL, d]) == 1 then "CALL " else "" endif ++
         if d == nb_days then "\n" else "" endif 
         | p in P, d in D
       ];