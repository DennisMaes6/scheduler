int: nb_weeks;
int: nb_days;
set of int: D;                % days

enum ShiftType = {
  O,
  JA_NW,
  SA_EW,
  JA_E,
  JA_WH,
  SA_WH,
  TS,
  C
};

set of ShiftType: S;          % shift types
set of ShiftType: S_prime;    % shift types - excluding free
set of ShiftType: S_week;     % week shift types
set of ShiftType: S_weekend;  % weekend shift types
set of ShiftType: S_weekday;  % weekday shift types

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


array[1..7, S] of int: R;               % required personnel per day per shift type                                        
array[S] of set of PersonnelType: Q;    % qualifications                                             
array[1..nb_personnel] of set of D: F;  % granted holidays per employee 
 
                                                                                               
% decision variables X[p, s, d] == 1 -> employee p is assigned shift of type s on day d
array[1..nb_personnel, S, D] of var 0..1: X;


%% CONSTRAINTS


% HC1 - shift requirements met for every shift on every day
constraint forall(s in S_prime, d in D)((sum(p in P)(X[p, s, d])) = R[1 + (d mod 7), s]);

% HC2 - only qualified personnel allowed
constraint forall(s in S_prime, p in P, d in D)(X[p, s, d] = 1 -> T[p] in Q[s]);


% HC3 - exaclty one assignment per personnel member on every day
constraint forall(p in P, d in D)(sum(s in S)(X[p, s, d]) = 1);

% HC4 - same person for week shifts
constraint forall(w in 1..nb_weeks, s in S_week)(exists(p in P)(forall(d in 7*w-6..7*w)(X[p, s, d] = 1)));

% HC5 - same person for weekend shifts
constraint forall(w in 1..nb_weeks, s in S_weekend, p in P)(X[p, s, 7*w-2] = X[p, s, 7*w-1]);


% HC6 - no assignments on free days
constraint forall(s in S_prime, p in P, d in D)(X[p, s, d] = 1 -> not (d in F[p]));


% HC7 - no consecutive week assignements
constraint forall(p in P, s in S_week, w in 2..nb_weeks)(
             exists(s_prev in S_week)(X[p, s_prev, 7*w - 8] = 1) -> X[p, s, 7*w - 6] = 0
           );
           
                           
% HC8 - no consecutive weekday assignements
constraint forall(p in P, s in S_weekday, d in 2..nb_days)(
             exists(s_prev in S_weekday)(X[p, s_prev, d - 1] = 1) -> X[p, s, d] = 0
           );
           

% optimization objective: fairness
array[S_prime] of var 0..1.0: fairness_per_shift_type;
constraint forall(s in S_prime)(fairness_per_shift_type[s] == max(p in P where T[p] in Q[s])(workload[s, p]) - min(p in P where T[p] in Q[s])(workload[s, p]));

array[S_prime, P] of var 0..1.0: workload;
constraint forall(s in S_prime, p in P)(workload[s, p] == (sum(d in D)(X[p, s, d])) / (card(D) - card(F[p])));

var 0..1.0: fairness_score;
constraint fairness_score == max(fairness_per_shift_type);
constraint fairness_score < 0.35;


% optimization objective: balance

array[P, D] of var 0..nb_days: a;
constraint forall(p in P, d in D)(
             if d = 1
             then 
               if sum(s in S_prime)(X[p, s, d]) = 1
               then a[p, d] = 0
               else a[p, d] = 8
               endif
             else 
               if sum(s in S_prime)(X[p, s, d]) = 0 
               then a[p, d] = 1 + a[p, d-1]
               else a[p, d] = 0
               endif
             endif);
           
array[P, D] of var 0..nb_days: z;
constraint forall(p in P, d in D)(
             if d = nb_days
             then  z[p, d] = nb_days
             else 
               if sum(s in S_prime)(X[p, s, d]) = 0 /\ sum(s in S_prime)(X[p, s, d+1]) = 1
               then z[p, d] = a[p, d]
               else z[p, d] = nb_days
               endif
             endif
           );
           
constraint min(z) >= 5;

solve satisfy;

output 
       ["day:" ++ show(d) ++ " " | d in D] ++ ["\n"] ++
       [
         if d == 1 then "assistant:" ++ show(p) ++ " " ++ "type:" ++ show(T[p]) ++ " " else "" endif ++
         if fix(X[p, O, d])     == 1 then "O " else "" endif ++
         if fix(X[p, JA_NW, d]) == 1 then "JA_NW " else "" endif ++
         if fix(X[p, SA_EW, d]) == 1 then "SA_EW " else "" endif ++
         if fix(X[p, JA_E, d])  == 1 then "JA_E " else "" endif ++
         if fix(X[p, JA_WH, d]) == 1 then "JA_WH " else "" endif ++
         if fix(X[p, SA_WH, d]) == 1 then "SA_WH " else "" endif ++
         if fix(X[p, TS, d])    == 1 then "TS " else "" endif ++
         if fix(X[p, C, d])     == 1 then "C " else "" endif ++
         if d == nb_days then "\n" else "" endif 
         | p in P, d in D
       ];