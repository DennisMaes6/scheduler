include "all_equal.mzn";

%%%%%%%%%%%%%
% VARAIBLES %
%%%%%%%%%%%%%

% days
int: nb_weeks;
int: nb_days = 7 * nb_weeks;
set of int: D = 1..nb_days;
set of int: H;                                     % holidays


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

array[ShiftType] of int: max_buffer;
array[ShiftType] of int: max_assignments = [
  ceil(nb_days / sum(p in P where T[p] in Q[s])(1)) + max_buffer[s] | s in ShiftType
];
array[ShiftType] of float: shift_workload;


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
set of 1..nb_personnel: P = 1..nb_personnel;            % employees
array[1..nb_personnel] of PersonnelType: T;             % mapping from employee to personnel type
array[1..nb_personnel] of set of 1..nb_days: F;         % granted holidays per employee 

% coverage requirements 
array[ShiftType, 1..7] of int: R = [|1, 1, 1, 1, 1, 1, 1 |
                                     0, 2, 2, 0, 0, 0, 0 |
                                     1, 1, 1, 1, 1, 1, 1 |
                                     0, 2, 2, 0, 0, 0, 0 |
                                     0, 1, 1, 0, 0, 0, 0 |
                                     1, 1, 1, 1, 1, 1, 1 |
                                     0, 0, 0, 0, 0, 0, 0 |];

function int: requirements(ShiftType: s, D: d) =
  if s in S_weekend /\ d in H
  then R[s, 2]
  else R[s, ((d - 1) mod 7) + 1]
  endif;

                            
% qualifications                                                                                                                                       
array[ShiftType] of set of PersonnelType: Q = [{JA},                            % JANW
                                               {JA, JA_F},                      % JAWH
                                               {SA, SA_NEO},                    % SAEW
                                               {SA, SA_F, SA_NEO, SA_F_NEO},    % SAWH
                                               {SA_NEO, SA_F_NEO},              % TSPT
                                               {SA, SA_F, SA_NEO, SA_F_NEO},    % CALL
                                               PersonnelType];                  % FREE
                                                                                          
% model variables
int: min_balance;

                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 
% decision variables X[p, s, d] = 1 -> employee p is assigned shift of type s on day d
array[1..nb_personnel, ShiftType, D] of var 0..1: X;

%%%%%%%%%%%%%%%
% CONSTRAINTS %
%%%%%%%%%%%%%%%


% HC1 - shift requirements met for every shift on every day
constraint forall(s in S, d in D)((sum(p in P)(X[p, s, d])) = 1);

% HC2 - only qualified personnel allowed
constraint forall(s in S, p in P, d in D)(X[p, s, d] = 1 -> T[p] in Q[s]);

% HC3 - exaclty one assignment per personnel member on every week
constraint forall(p in P, d in D)(sum(s in S_full)(X[p, s, d]) = 1);

% HC4 - same person for week shifts
constraint forall(w in 1..nb_weeks, s in S_week, p in P)(all_equal(X[p, s, 7*w-6..7*w]));

% HC5 - same person for weekend shifts
constraint forall(w in 1..nb_weeks, s in S_weekend, p in P)(X[p, s, 7*w-5] = X[p, s, 7*w-4]);

% HC6 - no assignments on free days
constraint forall(s in S, p in P, d in D)(X[p, s, d] = 1 -> not (d in F[p]));

% HC7 - no consecutive assignements of different shift types
constraint forall(s in S, p in P, d in 2..nb_days)(
            exists(s_prev in S where s_prev != s)(X[p, s_prev, d-1] = 1) -> X[p, s, d] = 0
           ); 

% HC7 - no consecutive week assignements
constraint forall(s in S_week, p in P, w in 2..nb_weeks)(
            exists(s_prev in S_week)(X[p, s_prev, 7*(w-1)] = 1) -> X[p, s, 7*w] = 0
           ); 
           
% HC8 - respect max assignments per shift type
constraint forall(p in P, s in S)(sum(d in D)(X[p, s, d]) <= max_assignments[s]);    

% optimization objective: fairness
array[P] of var float: workload;
constraint forall(p in P)(workload[p] = (sum(d in D, s in S)(X[p, s, d] * shift_workload[s]) / (nb_days - card(F[p]))));

var float: fairness_score;
constraint fairness_score = max(workload) - min(workload);


% optimization objective: balance
array[P, D] of var 0..nb_days: a;
constraint forall(p in P, d in D)(
  if d = 1
  then a[p, d] = 1 - sum(s in S)(X[p, s, d])
  else if sum(s in S)(X[p, s, d]) = 1
       then a[p, d] = 0
       else a[p, d] = 1 + a[p, d-1]
       endif
  endif
);


predicate after_first(P: p, D: d) =
  exists(d_prev in D, s in S)(d > d_prev /\ X[p, s, d_prev] = 1);
           
array[P, D] of var D: z;
constraint forall(p in P, d in D)(
  if d = nb_days
  then z[p, d] = nb_days
  else if after_first(p, d)
      then if (a[p, d] != 0 /\ a[p, d+1] == 0)
            then z[p, d] = a[p, d]
            else z[p, d] = nb_days
            endif
      else z[p, d] = nb_days
      endif
  endif
);
           
constraint min(z) >= min_balance;

solve minimize fairness_score;

output
       ["shift_types: "] ++ [show(s) ++ " " | s in S] ++ ["\n"] ++
       ["min_balance:"] ++ [show(min_balance)] ++ ["\n"] ++
       ["fairness_score:" ++ show(fairness_score)++"\n"] ++ 
       ["nb_days:"] ++ [show(nb_days)] ++ ["\n"] ++
       [
         if d == 1 then 
             "assistant:" ++ show(p) ++ " " 
          ++ "type:" ++ show(T[p]) ++ " " 
          ++ "workload:" ++ show(workload[p]) ++ " " 
         else "" endif ++

         if fix(X[p, FREE, d]) == 1 then "FREE " else "" endif ++
         if fix(X[p, JANW, d]) == 1 then "JANW " else "" endif ++
         if fix(X[p, SAEW, d]) == 1 then "SAEW " else "" endif ++
         if fix(X[p, JAWH, d]) == 1 then "JAWH " else "" endif ++
         if fix(X[p, SAWH, d]) == 1 then "SAWH " else "" endif ++
         if fix(X[p, TSPT, d]) == 1 then "TSPT " else "" endif ++
         if fix(X[p, CALL, d]) == 1 then "CALL " else "" endif ++
         if d == nb_days then "\n" else "" endif 
         | p in P, d in D
       ];