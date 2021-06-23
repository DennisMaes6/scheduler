include "all_equal.mzn";

%%%%%%%%%%%%%
% VARIABLES %
%%%%%%%%%%%%%

% days
int: nb_weeks;
int: nb_days = 7 * nb_weeks;
set of int: D = 1..nb_days;
set of int: W = 1..nb_weeks;
set of int: H;                 % holidays
int: nb_holidays; 


% shift types
enum ShiftType = {
  JANW,
  JAWE,
  JAHO,
  SAEW,
  SAWE,
  SAHO,
  TPWE,
  TPHO,
  CALL,
  FREE
};

set of ShiftType: S = {JANW, SAEW, JAWE, JAHO, SAWE, SAHO, TPWE, TPHO, CALL};   
set of ShiftType: S_full = S union {FREE};   
set of ShiftType: S_week = S intersect {JANW, SAEW, CALL};    
set of ShiftType: S_weekend = S intersect {JAWE, SAWE, TPWE};
set of ShiftType: S_holiday = S intersect {JAHO, SAHO, TPHO};

array[ShiftType] of int: max_buffer;
array[ShiftType] of int: max_assignments = [
  if s in S_holiday 
    then ceil(nb_holidays / sum(p in P where T[p] in Q[s])(1)) + max_buffer[s]
    else ceil(nb_weeks / sum(p in P where T[p] in Q[s])(1)) + max_buffer[s]
  endif 
  | s in ShiftType
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
array[1..nb_personnel] of set of D: F;                  % granted holidays per employee 

% coverage requirements 
array[ShiftType] of int: R = [1, 2, 2, 1, 2, 2, 1, 1, 1, 0];
                            
% qualifications                                                                                                                                       
array[ShiftType] of set of PersonnelType: Q = [{JA},                            % JANW
                                               {JA, JA_F},                      % JAWE
                                               {JA, JA_F},                      % JAHO
                                               {SA, SA_NEO},                    % SAEW
                                               {SA, SA_F, SA_NEO, SA_F_NEO},    % SAWE
                                               {SA, SA_F, SA_NEO, SA_F_NEO},    % SAHO
                                               {SA_NEO, SA_F_NEO},              % TPWE
                                               {SA_NEO, SA_F_NEO},              % TPHO
                                               {SA, SA_F, SA_NEO, SA_F_NEO},    % CALL
                                               PersonnelType];                  % FREE

function int: day_to_week(D: d) =
  ceil((d-0.5) / 7);

function ShiftType: weekend_to_holiday_shift(S_weekend: s) =
  if s = JAWE then JAHO else
    if s = SAWE then SAHO else
      if s = TPWE then TPHO else s endif % default case should not be reached!
    endif
  endif;

function bool: is_weekend(D: d) = 
  d mod 7 = 2 \/ d mod 7 = 3;


                                                                                          
% model variables
int: min_balance;
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                
% decision variables X[p, s, d] == 1 -> employee p is assigned shift of type s on day d
array[1..nb_personnel, ShiftType, W] of var 0..1: X;
array[1..nb_personnel, ShiftType, D] of var 0..1: Xh;

%%%%%%%%%%%%%%%
% CONSTRAINTS %
%%%%%%%%%%%%%%%

% HC1 - shift requirements met for every shift on every day
constraint forall(s in (S_week union S_weekend), w in W)((sum(p in P)(X[p, s, w])) = R[s]);
constraint forall(s in S_holiday, d in H)((sum(p in P)(Xh[p, s, d])) = R[s]);

% HC2 - only qualified personnel allowed
constraint forall(s in (S_week union S_weekend), p in P, w in W)(X[p, s, w] = 1 -> T[p] in Q[s]);
constraint forall(s in S_holiday, p in P, d in H)(Xh[p, s, d] = 1 -> T[p] in Q[s]);

% HC3 - exaclty one assignment per personnel member on every week
constraint forall(p in P, w in W)(sum(s in (S_week union S_weekend union {FREE}))(X[p, s, w]) = 1);
constraint forall(p in P, d in H)(sum(s in S_holiday)(Xh[p, s, d]) <= 1);
constraint forall(p in P, d in H)(sum(s in S_holiday)(Xh[p, s, d]) = 1 -> sum(s in S_week)(X[p, s, day_to_week(d)]) = 0);

% HC6 - no assignments on free days
constraint forall(s in S_week, p in P, w in W)(X[p, s, w] = 1 -> (7*w-6..7*w intersect F[p] = {}));
constraint forall(s in S_weekend, p in P, w in W)(X[p, s, w] = 1 -> (7*w-5..7*w-4 intersect F[p] = {}));
constraint forall(s in S_holiday, p in P, d in H)(Xh[p, s, d] = 1 -> not(d in F[p]));

% HC7 - no consecutive assignements
constraint forall(p in P, w in 1..nb_weeks-1, s in (S_week))(
             exists(s_prev in (S_week))(X[p, s_prev, w] = 1) -> X[p, s, w+1] = 0
           );
constraint forall(p in P, h in H where h mod 7 = 1 /\ h > 1)(
  exists(sh in S_holiday)(Xh[p, sh, h] = 1) -> sum(s in S_week)(X[p, s, day_to_week(h-1)]) = 0
);
constraint forall(p in P, h in H where h mod 7 = 0 /\ h < 7*nb_weeks)(
  exists(sh in S_holiday)(Xh[p, sh, h] = 1) -> sum(s in S_week)(X[p, s, day_to_week(h+1)]) = 0
);
constraint forall(p in P, h in H where h mod 7 = 1 /\ not exists(h_after in H)(h + 1 = h_after))(
  exists(sh in S_holiday)(Xh[p, sh, h] = 1) -> sum(s in S_weekend)(X[p, s, day_to_week(h)]) = 0
);
constraint forall(p in P, h in H where h mod 7 = 4 /\ not exists(h_before in H)(h - 1 = h_before))(
  exists(sh in S_holiday)(Xh[p, sh, h] = 1) -> sum(s in S_weekend)(X[p, s, day_to_week(h)]) = 0
);
           
% HC8 - respect max assignments per shift type
constraint forall(p in P, s in (S_week union S_weekend))(sum(w in W)(X[p, s, w]) <= max_assignments[s]);
constraint forall(p in P, s in S_holiday)(sum(h in H)(Xh[p, s, h]) <= max_assignments[s]);



% consecutive holidays same person
constraint forall(h1 in H, h2 in H where h1 + 1 = h2)(
  forall(p in P, s in S)(Xh[p, s, h1] = 1 <-> Xh[p, s, h2] = 1)
);

% holidays in weekend same person
constraint forall(h in H where is_weekend(h))(
  forall(p in P, s in S_weekend)(X[p, s, day_to_week(h)] = 1 -> Xh[p, weekend_to_holiday_shift(s), h] = 1)
);


% optimization objective: fairness
array[P] of var float: workload;
constraint forall(p in P)(workload[p] = 
  if card(F[p]) == nb_days
  then 0
  else (sum(w in W, s in (S_week union S_weekend))(X[p, s, w] * shift_workload[s])
        + sum(d in H, s in S_holiday)(Xh[p, s, d] * shift_workload[s]))/ (nb_days - card(F[p]))
  endif
);

var float: fairness_score;
constraint fairness_score = max(workload) - min(workload);


% balance
/*
array[P, 1..nb_days] of var 0..nb_days: a;
constraint forall(p in P, w in W)(
             if w = 1
             then if sum(s in S_week)(X[p, s, w]) = 1
                  then forall(i in 7*w-6..7*w)(a[p, i] = 0)
                  else if sum(s in S_weekend)(X[p, s, w]) = 1
                       then  a[p, 1] = if 1 in H /\ sum(sh in S_holiday)(Xh[p, sh, 1]) = 1 then 0 else 1           endif 
                          /\ a[p, 2] = 0
                          /\ a[p, 3] = 0 
                          /\ a[p, 4] = if 4 in H /\ sum(sh in S_holiday)(Xh[p, sh, 4]) = 1 then 0 else 1 + a[p, 3] endif
                          /\ a[p, 5] = if 5 in H /\ sum(sh in S_holiday)(Xh[p, sh, 5]) = 1 then 0 else 1 + a[p, 4] endif
                          /\ a[p, 6] = if 6 in H /\ sum(sh in S_holiday)(Xh[p, sh, 6]) = 1 then 0 else 1 + a[p, 5] endif
                          /\ a[p, 7] = if 7 in H /\ sum(sh in S_holiday)(Xh[p, sh, 7]) = 1 then 0 else 1 + a[p, 6] endif
                       else  a[p, 1] = if 1 in H /\ sum(sh in S_holiday)(Xh[p, sh, 1]) = 1 then 0 else 1           endif 
                          /\ a[p, 2] = if 2 in H /\ sum(sh in S_holiday)(Xh[p, sh, 2]) = 1 then 0 else 1 + a[p, 1] endif
                          /\ a[p, 3] = if 3 in H /\ sum(sh in S_holiday)(Xh[p, sh, 3]) = 1 then 0 else 1 + a[p, 2] endif
                          /\ a[p, 4] = if 4 in H /\ sum(sh in S_holiday)(Xh[p, sh, 4]) = 1 then 0 else 1 + a[p, 3] endif
                          /\ a[p, 5] = if 5 in H /\ sum(sh in S_holiday)(Xh[p, sh, 5]) = 1 then 0 else 1 + a[p, 4] endif
                          /\ a[p, 6] = if 6 in H /\ sum(sh in S_holiday)(Xh[p, sh, 6]) = 1 then 0 else 1 + a[p, 5] endif
                          /\ a[p, 7] = if 7 in H /\ sum(sh in S_holiday)(Xh[p, sh, 7]) = 1 then 0 else 1 + a[p, 6] endif
                       endif
                  endif
             else 
               if sum(s in S_week)(X[p, s, w]) = 1 
               then forall(i in 7*w-6..7*w)(a[p, i] = 0)
               else if sum(s in S_weekend)(X[p, s, w]) = 1
                    then  a[p, 7*w-6] = if 7*w-6 in H /\ sum(sh in S_holiday)(Xh[p, sh, 7*w-6]) = 1 then 0 else 1 + a[p, 7*w-7] endif
                       /\ a[p, 7*w-5] = 0
                       /\ a[p, 7*w-4] = 0 
                       /\ a[p, 7*w-3] = if 7*w-3 in H /\ sum(sh in S_holiday)(Xh[p, sh, 7*w-3]) = 1 then 0 else 1 + a[p, 7*w-4] endif
                       /\ a[p, 7*w-2] = if 7*w-2 in H /\ sum(sh in S_holiday)(Xh[p, sh, 7*w-2]) = 1 then 0 else 1 + a[p, 7*w-3] endif
                       /\ a[p, 7*w-1] = if 7*w-1 in H /\ sum(sh in S_holiday)(Xh[p, sh, 7*w-1]) = 1 then 0 else 1 + a[p, 7*w-2] endif
                       /\ a[p, 7*w  ] = if 7*w   in H /\ sum(sh in S_holiday)(Xh[p, sh, 7*w  ]) = 1 then 0 else 1 + a[p, 7*w-1] endif
                    else forall(i in 7*w-6..7*w)(a[p, i] = if i in H /\ sum(sh in S_holiday)(Xh[p, sh, i]) = 1 then 0 else 1 + a[p, i-1] endif)
                    endif
               endif
             endif);
*/

array[P, 1..nb_days] of var 0..nb_days: a;
constraint forall(p in P, w in W)(
             if w = 1
             then if sum(s in S_week)(X[p, s, w]) = 1
                  then forall(i in 7*w-6..7*w)(a[p, i] = 0)
                  else if sum(s in S_weekend)(X[p, s, w]) = 1
                       then  a[p, 1] = 1
                          /\ a[p, 2] = 0
                          /\ a[p, 3] = 0 
                          /\ a[p, 4] = 1 + a[p, 3]
                          /\ a[p, 5] = 1 + a[p, 4]
                          /\ a[p, 6] = 1 + a[p, 5]
                          /\ a[p, 7] = 1 + a[p, 6]
                       else  a[p, 1] = 1           
                          /\ a[p, 2] = 1 + a[p, 1]
                          /\ a[p, 3] = 1 + a[p, 2]
                          /\ a[p, 4] = 1 + a[p, 3]
                          /\ a[p, 5] = 1 + a[p, 4]
                          /\ a[p, 6] = 1 + a[p, 5]
                          /\ a[p, 7] = 1 + a[p, 6]
                       endif
                  endif
             else 
               if sum(s in S_week)(X[p, s, w]) = 1 
               then forall(i in 7*w-6..7*w)(a[p, i] = 0)
               else if sum(s in S_weekend)(X[p, s, w]) = 1
                    then  a[p, 7*w-6] = 1 + a[p, 7*w-7]
                       /\ a[p, 7*w-5] = 0
                       /\ a[p, 7*w-4] = 0 
                       /\ a[p, 7*w-3] = 1 + a[p, 7*w-4]
                       /\ a[p, 7*w-2] = 1 + a[p, 7*w-3]
                       /\ a[p, 7*w-1] = 1 + a[p, 7*w-2]
                       /\ a[p, 7*w  ] = 1 + a[p, 7*w-1]
                    else forall(i in 7*w-6..7*w)(a[p, i] = 1 + a[p, i-1])
                    endif
               endif
             endif);




predicate after_first(1..nb_personnel: p, D: d) =
  exists(h in H, s in S_holiday)(d > h /\ Xh[p, s, h] = 1) \/
  exists(w in W, s in S_week)(d > 7*w /\ X[p, s, w] = 1) \/
  exists(w in W, s in S_weekend)(d > 7*(w-1)+3 /\ X[p, s, w] = 1);
           
array[P, D] of var 0..nb_days: z;
constraint forall(p in P, d in D)(
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
       %[if d = 1 then "\n" else " " endif ++ show(fix(a[p, d])) | p in P, d in D] ++ ["\n\n"] ++
       %£[if d = 1 then "\n" else " " endif ++ show(fix(z[p, d])) | p in P, d in D] ++
       ["balance_score:"] ++ [show(min(z))] ++ ["\n"] ++
       ["fairness_score:" ++ show(fairness_score)++"\n"] ++ 
       [
         if w == 1 then 
             "assistant_id:" ++ show(personnel_id[p]) ++ " " 
          ++ "workload:" ++ if fix(workload[p]) <= 0.0 then "0.0" else show(workload[p]) endif ++ " " 
         else "" endif ++
         if fix(X[p, FREE, w]) == 1 
         then if 7*w-6..7*w intersect H = {} 
              then "FREE FREE FREE FREE FREE FREE FREE "
              else concat([
                if fix(Xh[p, JAHO, d]) == 1 then "JAHO " else
                  if fix(Xh[p, SAHO, d]) == 1 then "SAHO " else
                    if fix(Xh[p, TPHO, d]) == 1 then "TPHO " else
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
         if fix(X[p, JAWE, w]) == 1 
         then if 7*w-6..7*w intersect H = {} 
          then "FREE JAWE JAWE FREE FREE FREE FREE "
            else concat([
              if fix(Xh[p, JAHO, d]) == 1 then "JAHO " else
                if fix(Xh[p, SAHO, d]) == 1 then "SAHO " else
                  if fix(Xh[p, TPHO, d]) == 1 then "TPHO " else
                    if d mod 7 = 2 \/ d mod 7 = 3
                    then "JAWE "
                    else "FREE "
                    endif
                  endif
                endif
              endif
              | d in 7*w-6..7*w])
            endif
         else "" 
         endif ++
         if fix(X[p, SAWE, w]) == 1 
          then if 7*w-6..7*w intersect H = {} 
            then "FREE SAWE SAWE FREE FREE FREE FREE "
              else concat([
                if fix(Xh[p, JAHO, d]) == 1 then "JAHO " else
                  if fix(Xh[p, SAHO, d]) == 1 then "SAHO " else
                    if fix(Xh[p, TPHO, d]) == 1 then "TPHO " else
                      if d mod 7 = 2 \/ d mod 7 = 3
                      then "SAWE "
                      else "FREE "
                      endif
                    endif
                  endif
                endif
                | d in 7*w-6..7*w])
              endif
         else "" 
         endif ++
         if fix(X[p, TPWE, w]) == 1 
          then if 7*w-6..7*w intersect H = {} 
            then "FREE TPWE TPWE FREE FREE FREE FREE "
              else concat([
                if fix(Xh[p, JAHO, d]) == 1 then "JAHO " else
                  if fix(Xh[p, SAHO, d]) == 1 then "SAHO " else
                    if fix(Xh[p, TPHO, d]) == 1 then "TPHO " else
                      if d mod 7 = 2 \/ d mod 7 = 3
                      then "TPWE "
                      else "FREE "
                      endif
                    endif
                  endif
                endif
                | d in 7*w-6..7*w])
              endif
         else "" 
         endif ++
         if fix(X[p, CALL, w]) == 1 then "CALL CALL CALL CALL CALL CALL CALL " else "" endif ++
         if w == nb_weeks then "\n" else "" endif
         | p in P, w in W
       ];