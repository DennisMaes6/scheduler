include "all_equal.mzn";

%%%%%%%%%%%%%
% VARAIBLES %
%%%%%%%%%%%%%

% days
int: nb_weeks;
int: nb_days = 7 * nb_weeks;
set of int: W = 1..nb_weeks;


% shift types
enum ShiftType = {
  JANW,
  JAWH,
  SAEW,
  SAWH,
  TSPT,
  CALL,
  FREE
};

set of ShiftType: S;     
set of ShiftType: S_full = S union {FREE};   
set of ShiftType: S_week = S intersect {JANW, SAEW, CALL};    
set of ShiftType: S_weekend = S intersect {JAWH, SAWH, TSPT};


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
array[1..nb_personnel] of set of 1..nb_days: F;         % granted holidays per employee 

% coverage requirements 
array[ShiftType] of int: R = [1, 2, 1, 2, 1, 1, 0];
                            
% qualifications                                                                                                                                       
array[ShiftType] of set of PersonnelType: Q = [{JA},                            % JANW
                                               {JA, JA_F},                      % JAWH
                                               {SA, SA_NEO},                    % SAEW
                                               {SA, SA_F, SA_NEO, SA_F_NEO},    % SAWH
                                               {SA_NEO, SA_F_NEO},              % TSPT
                                               {SA, SA_F, SA_NEO, SA_F_NEO},    % CALL
                                               PersonnelType];                  % FREE
                                                                                          
% model variables
array[ShiftType] of float: fairness_weight;
set of ShiftType: S_balance;
int: min_balance;

                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 
% decision variables X[p, s, d] == 1 -> employee p is assigned shift of type s on day d
array[1..nb_personnel, ShiftType, 1..nb_weeks] of var 0..1: X;

%%%%%%%%%%%%%%%
% CONSTRAINTS %
%%%%%%%%%%%%%%%


% HC1 - shift requirements met for every shift on every day
constraint forall(s in S, w in W)((sum(p in P)(X[p, s, w])) = R[s]);

% HC2 - only qualified personnel allowed
constraint forall(s in S, p in P, w in W)(X[p, s, w] = 1 -> T[p] in Q[s]);

% HC3 - exaclty one assignment per personnel member on every week
constraint forall(p in P, w in W)(sum(s in S_full)(X[p, s, w]) = 1);

% HC6 - no assignments on free days
constraint forall(s in S, p in P, w in W)(X[p, s, w] = 1 -> (7*w-6..7*w intersect F[p] = {}));

% HC7 - no consecutive week assignements
constraint forall(p in P, s in S_week, w in 2..nb_weeks)(
             exists(s_prev in S)(X[p, s_prev, w-1] = 1) -> X[p, s, w] = 0
           );       


% optimization objective: fairness
array[ShiftType] of var 0..1.0: fairness_per_shift_type;
constraint forall(s in S)(fairness_per_shift_type[s] == max(p in P where T[p] in Q[s])(workload[s, p]) - min(p in P where T[p] in Q[s])(workload[s, p]));

function var int: workload_shift_type(var ShiftType: st) =
  if st in S_week then 7 else
    if st in S_weekend then 7 else 0 endif
  endif;

array[ShiftType, P] of var 0..1.0: workload;
constraint forall(s in S, p in P)(workload[s, p] == ((sum(w in W)(X[p, s, w])) * workload_shift_type(s)) / (nb_days - card(F[p])));

array[ShiftType] of var float: weighted_fairness;
constraint forall(s in S)(weighted_fairness[s] = fairness_per_shift_type[s] * fairness_weight[s]);

var 0..1.0: fairness_score;
constraint fairness_score == max(s in S)(weighted_fairness[s]);


% optimization objective: balance
array[P, 1..nb_days] of var 0..nb_days: a;
constraint forall(p in P, w in W)(
             if w = 1
             then 
               if sum(s in S_week intersect S_balance)(X[p, s, w]) = 1
               then forall(i in 7*w-6..7*w)(a[p, i] = 0)
               else if sum(s in S_weekend intersect S_balance)(X[p, s, w]) = 1
                    then  a[p, 7*w-6] = 0 
                       /\ a[p, 7*w-5] = 0
                       /\ a[p, 7*w-4] = 0 
                       /\ a[p, 7*w-3] = 1 
                       /\ a[p, 7*w-2] = 2 
                       /\ a[p, 7*w-1] = 3 
                       /\ a[p, 7*w]   = 4 
                    else forall(i in 7*w-6..7*w)(a[p, i] = i)
                    endif
               endif
             else 
               if sum(s in S_week intersect S_balance)(X[p, s, w]) = 1 
               then forall(i in 7*w-6..7*w)(a[p, i] = 0)
               else if sum(s in S_weekend intersect S_balance)(X[p, s, w]) = 1
                    then  a[p, 7*w-6] = 1 + a[p, 7*w-7]
                       /\ a[p, 7*w-5] = 0
                       /\ a[p, 7*w-4] = 0 
                       /\ a[p, 7*w-3] = 1
                       /\ a[p, 7*w-2] = 2 
                       /\ a[p, 7*w-1] = 3 
                       /\ a[p, 7*w]   = 4 
                    else forall(i in 7*w-6..7*w)(a[p, i] = 1 + a[p, i-1])
                    endif
               endif
             endif);
           
array[P, 1..nb_days] of var 0..nb_days: z;
constraint forall(p in P, d in 1..nb_days)(
             if d = nb_days
             then  z[p, d] = nb_days
             else 
               if (a[p, d] != 0 /\ a[p, d+1] == 0)
               then z[p, d] = a[p, d]
               else z[p, d] = nb_days
               endif
             endif
           );
           
constraint min(z) >= min_balance;

solve minimize fairness_score;

 output  
       ["balance_shifts:"] ++ [show(s) ++ " " | s in S_balance] ++ ["\n"] ++
       ["min_balance:"] ++ [show(min_balance)] ++ ["\n"] ++
       ["nb_days:"] ++ [show(nb_days)] ++ ["\n"] ++
       ["fairness_per_shift_type:"] ++ [show(s) ++ "=" ++ show(weighted_fairness[s]) ++ " " | s in S] ++ ["\n"] ++
       [
         if w == 1 then "assistant:" ++ show(p) ++ " " ++ "type:" ++ show(T[p]) ++ " " else "" endif ++
         if fix(X[p, FREE, w]) == 1 then "FREE FREE FREE FREE FREE FREE FREE " else "" endif ++
         if fix(X[p, JANW, w]) == 1 then "JANW JANW JANW JANW JANW JANW JANW " else "" endif ++
         if fix(X[p, SAEW, w]) == 1 then "SAEW SAEW SAEW SAEW SAEW SAEW SAEW " else "" endif ++
         if fix(X[p, JAWH, w]) == 1 then "FREE JAWH JAWH FREE FREE FREE FREE " else "" endif ++
         if fix(X[p, SAWH, w]) == 1 then "FREE SAWH SAWH FREE FREE FREE FREE " else "" endif ++
         if fix(X[p, TSPT, w]) == 1 then "FREE TSPT TSPT FREE FREE FREE FREE " else "" endif ++
         if fix(X[p, CALL, w]) == 1 then "CALL CALL CALL CALL CALL CALL CALL " else "" endif ++
         if w == nb_weeks then "\n" else "" endif 
         | p in P, w in W
       ];