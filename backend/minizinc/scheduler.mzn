include "all_equal.mzn";

%%%%%%%%%%%%%
% VARIABLES %
%%%%%%%%%%%%%

% days
int: nb_weeks;
int: nb_days = 7 * nb_weeks;
set of int: D = 1..nb_days;
set of int: W = 1..nb_weeks;
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

set of ShiftType: S = {JANW, SAEW, JAWH, SAWH, TSPT, CALL};   
set of ShiftType: S_full = S union {FREE};   
set of ShiftType: S_week = S intersect {JANW, SAEW, CALL};    
set of ShiftType: S_weekend = S intersect {JAWH, SAWH, TSPT};

array[ShiftType] of int: max_buffer;
array[ShiftType] of int: max_assignments = [
  ceil(nb_weeks / sum(p in P where T[p] in Q[s])(1)) + max_buffer[s] | s in ShiftType
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
array[1..nb_personnel] of int: personnel_id;            % mapping from employee to external id
array[1..nb_personnel] of PersonnelType: T;             % mapping from employee to personnel type
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

function int: day_to_week(D: d) =
  ceil((d-0.5) / 7);

constraint day_to_week(2) = 1;
                                                                                          
% model variables
int: min_balance;
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                
% decision variables X[p, s, d] == 1 -> employee p is assigned shift of type s on day d
array[1..nb_personnel, ShiftType, W] of var 0..1: X;
array[1..nb_personnel, ShiftType, D] of var 0..1: Xh;

%%%%%%%%%%%%%%%
% CONSTRAINTS %
%%%%%%%%%%%%%%%

% HC1 - shift requirements met for every shift on every day
constraint forall(s in S, w in W)((sum(p in P)(X[p, s, w])) = R[s]);
constraint forall(s in S_weekend, d in H)((sum(p in P)(Xh[p, s, d])) = R[s]);

% HC2 - only qualified personnel allowed
constraint forall(s in S, p in P, w in W)(X[p, s, w] = 1 -> T[p] in Q[s]);
constraint forall(s in S_weekend, p in P, d in H)(Xh[p, s, d] = 1 -> T[p] in Q[s]);

% HC3 - exaclty one assignment per personnel member on every week
constraint forall(p in P, w in W)(sum(s in S_full)(X[p, s, w]) = 1);
constraint forall(p in P, d in H)(sum(s in S_weekend)(Xh[p, s, d]) <= 1);
constraint forall(p in P, h in H)(sum(sh in S_weekend)(Xh[p, sh, h]) = 1 -> sum(s in S)(X[p, s, day_to_week(h)]) = 0);

% HC6 - no assignments on free days
constraint forall(s in S, p in P, w in W)(X[p, s, w] = 1 -> (7*w-6..7*w intersect F[p] = {}));
constraint forall(s in S_weekend, p in P, d in H)(Xh[p, s, d] = 1 -> not(d in F[p]));

% HC7 - no consecutive assignements
constraint forall(p in P, w in 1..nb_weeks-1, s in S)(
             exists(s_prev in S)(X[p, s_prev, w] = 1) -> X[p, s, w+1] = 0
           );
constraint forall(p in P, h in H where h mod 7 = 1 /\ h > 1)(
  exists(sh in S_weekend)(Xh[p, sh, h] = 1) -> sum(s in S_week)(X[p, s, day_to_week(h-1)]) = 0
);
constraint forall(p in P, h in H where h mod 7 = 0 /\ h < 7*nb_weeks)(
  exists(sh in S_weekend)(Xh[p, sh, h] = 1) -> sum(s in S_week)(X[p, s, day_to_week(h+1)]) = 0
);
           
% HC8 - respect max assignments per shift type
constraint forall(p in P, s in S)(sum(w in W)(X[p, s, w]) <= max_assignments[s]);    

% optimization objective: fairness
array[P] of var float: workload;
constraint forall(p in P)(workload[p] = (sum(w in W, s in S)(X[p, s, w] * shift_workload[s]) + sum(d in H, s in S_weekend)(Xh[p, s, d] * shift_workload[s] / 2))/ (nb_days - card(F[p])));

var float: fairness_score;
constraint fairness_score = max(workload) - min(workload);


% optimization objective: balance
array[P, 1..nb_days] of var 0..nb_days: a;
constraint forall(p in P, w in W)(
             if w = 1
             then if sum(s in S_week)(X[p, s, w]) = 1
                  then forall(i in 7*w-6..7*w)(a[p, i] = 0)
                  else if sum(s in S_weekend)(X[p, s, w]) = 1
                       then  a[p, 1] = if 1 in H /\ sum(sh in S_weekend)(Xh[p, sh, 1]) = 1 then 0 else 1           endif 
                          /\ a[p, 2] = 0
                          /\ a[p, 3] = 0 
                          /\ a[p, 4] = if 4 in H /\ sum(sh in S_weekend)(Xh[p, sh, 4]) = 1 then 0 else 1 + a[p, 3] endif
                          /\ a[p, 5] = if 5 in H /\ sum(sh in S_weekend)(Xh[p, sh, 5]) = 1 then 0 else 1 + a[p, 4] endif
                          /\ a[p, 6] = if 6 in H /\ sum(sh in S_weekend)(Xh[p, sh, 6]) = 1 then 0 else 1 + a[p, 5] endif
                          /\ a[p, 7] = if 7 in H /\ sum(sh in S_weekend)(Xh[p, sh, 7]) = 1 then 0 else 1 + a[p, 6] endif
                       else  a[p, 1] = if 1 in H /\ sum(sh in S_weekend)(Xh[p, sh, 1]) = 1 then 0 else 1           endif 
                          /\ a[p, 2] = if 2 in H /\ sum(sh in S_weekend)(Xh[p, sh, 2]) = 1 then 0 else 1 + a[p, 1] endif
                          /\ a[p, 3] = if 3 in H /\ sum(sh in S_weekend)(Xh[p, sh, 3]) = 1 then 0 else 1 + a[p, 2] endif
                          /\ a[p, 4] = if 4 in H /\ sum(sh in S_weekend)(Xh[p, sh, 4]) = 1 then 0 else 1 + a[p, 3] endif
                          /\ a[p, 5] = if 5 in H /\ sum(sh in S_weekend)(Xh[p, sh, 5]) = 1 then 0 else 1 + a[p, 4] endif
                          /\ a[p, 6] = if 6 in H /\ sum(sh in S_weekend)(Xh[p, sh, 6]) = 1 then 0 else 1 + a[p, 5] endif
                          /\ a[p, 7] = if 7 in H /\ sum(sh in S_weekend)(Xh[p, sh, 7]) = 1 then 0 else 1 + a[p, 6] endif
                       endif
                  endif
             else 
               if sum(s in S_week)(X[p, s, w]) = 1 
               then forall(i in 7*w-6..7*w)(a[p, i] = 0)
               else if sum(s in S_weekend)(X[p, s, w]) = 1
                    then  a[p, 7*w-6] = if 7*w-6 in H /\ sum(sh in S_weekend)(Xh[p, sh, 7*w-6]) = 1 then 0 else 1 + a[p, 7*w-7] endif
                       /\ a[p, 7*w-5] = 0
                       /\ a[p, 7*w-4] = 0 
                       /\ a[p, 7*w-3] = if 7*w-3 in H /\ sum(sh in S_weekend)(Xh[p, sh, 7*w-3]) = 1 then 0 else 1 + a[p, 7*w-4] endif
                       /\ a[p, 7*w-2] = if 7*w-2 in H /\ sum(sh in S_weekend)(Xh[p, sh, 7*w-2]) = 1 then 0 else 1 + a[p, 7*w-3] endif
                       /\ a[p, 7*w-1] = if 7*w-1 in H /\ sum(sh in S_weekend)(Xh[p, sh, 7*w-1]) = 1 then 0 else 1 + a[p, 7*w-2] endif
                       /\ a[p, 7*w  ] = if 7*w   in H /\ sum(sh in S_weekend)(Xh[p, sh, 7*w  ]) = 1 then 0 else 1 + a[p, 7*w-1] endif
                    else forall(i in 7*w-6..7*w)(a[p, i] = if i in H /\ sum(sh in S_weekend)(Xh[p, sh, i]) = 1 then 0 else 1 + a[p, i-1] endif)
                    endif
               endif
             endif);


predicate after_first(1..nb_personnel: p, 1..nb_days: d) =
  exists(h in H, s in S_weekend)(d > h /\ Xh[p, s, h] = 1) \/
  exists(w in W, s in S)(d > 7*w /\ X[p, s, w] = 1);
           
array[P, 1..nb_days] of var 0..nb_days: z;
constraint forall(p in P, d in 1..nb_days)(
             if d = nb_days
             then  z[p, d] = nb_days
             else
                if after_first(p, d)
                then 
                  if (a[p, d] != 0 /\ a[p, d+1] == 0)
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
       ["balance_score:"] ++ [show(min(z))] ++ ["\n"] ++
       ["fairness_score:" ++ show(fairness_score)++"\n"] ++ 
       [
         if w == 1 then 
             "assistant_id:" ++ show(personnel_id[p]) ++ " " 
          ++ "type:" ++ show(T[p]) ++ " " 
          ++ "workload:" ++ show(workload[p]) ++ " " 
         else "" endif ++
         if fix(X[p, FREE, w]) == 1 
         then if 7*w-6..7*w intersect H = {} 
              then "FREE FREE FREE FREE FREE FREE FREE "
              else concat([
                if fix(Xh[p, JAWH, d]) == 1 then "JAWH " else
                  if fix(Xh[p, SAWH, d]) == 1 then "SAWH " else
                    if fix(Xh[p, TSPT, d]) == 1 then "TSPT " else
                      "FREE "
                    endif
                  endif
                endif
                | d in 7*w-6..7*w])
              endif
         else "" 
         endif ++
         if fix(X[p, JANW, w]) == 1 then "JANW JANW JANW JANW JANW JANW JANW " else "" endif ++
         if fix(X[p, SAEW, w]) == 1 then "SAEW SAEW SAEW SAEW SAEW SAEW SAEW " else "" endif ++
         if fix(X[p, JAWH, w]) == 1 then "FREE JAWH JAWH FREE FREE FREE FREE " else "" endif ++
         if fix(X[p, SAWH, w]) == 1 then "FREE SAWH SAWH FREE FREE FREE FREE " else "" endif ++
         if fix(X[p, TSPT, w]) == 1 then "FREE TSPT TSPT FREE FREE FREE FREE " else "" endif ++
         if fix(X[p, CALL, w]) == 1 then "CALL CALL CALL CALL CALL CALL CALL " else "" endif ++
         if w == nb_weeks then "\n" else "" endif 
         | p in P, w in W
       ];