include "all_equal.mzn";

%%%%%%%%%%%%%
% VARAIBLES %
%%%%%%%%%%%%%

% days
int: nb_weeks;
int: nb_days = 7 * nb_weeks;
set of int: D = 1..nb_days;
set of D: H;   % holidays


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


% personnel
enum PersonnelType = {
  JA,
  JA_F
};

int: nb_personnel;
set of 1..nb_personnel: P = 1..nb_personnel;            % employees
array[1..nb_personnel] of int: personnel_id;            % mapping from employee to external id
array[1..nb_personnel] of set of 1..nb_days: F;         % granted holidays per employee 

% schedule
array[1..nb_personnel, 1..nb_days] of ShiftType: schedule;

% coverage requirements 
array[1..7] of int: R = [1, 0, 0, 1, 1, 1, 1];                         
                                                                                          
% model variables
int: min_balance;

% decision variable: X[p, d] == 1 => p is assigned a JAEV shift on day d
array[1..nb_personnel, 1..nb_days] of var 0..1: X;


% HC1 - only assignments on free days
constraint forall(p in P, d in D)(X[p, d] = 1 -> schedule[p, d] = FREE);

% HC2 - coverage requirements met
constraint forall(d in D)(sum(p in P)(X[p, d]) = if d in H
                                                 then 0 
                                                 else R[1 + ((d - 1) mod 7)]
                                                 endif);

% HC3 - no assignments on free days
constraint forall(p in P, d in D)(X[p, d] = 1 -> not (d in F[p]));

% HC4 - no consecutive jaev assignements
constraint forall(p in P, d in 2..nb_days)(X[p, d-1] = 1 -> X[p, d] = 0); 

% HC5 - no assignments three days before or after night weeks
constraint forall(p in P, d in 1..7)(X[p, d] = 1 -> schedule[p, d+3] != JANW); 
constraint forall(p in P, d in nb_days-6..nb_days)(X[p, d] = 1 -> schedule[p, d-3] != JANW); 
constraint forall(p in P, d in 8..nb_days-7)(X[p, d] = 1 -> schedule[p, d-3] != JANW /\ schedule[p, d+3] != JANW); 

% HC6 - no assignments one day before or after weekends
constraint forall(p in P)(X[p, 1] = 1 -> schedule[p, 2] != JAWE); 
constraint forall(p in P)(X[p, nb_days] = 1 -> schedule[p, nb_days-1] != JAWE); 
constraint forall(p in P, d in 2..nb_days-1)(X[p, d] = 1 -> schedule[p, d-1] != JAWE /\ schedule[p, d+1] != JAWE); 

% HC6 - no assignments one day before or after holidays
constraint forall(p in P)(X[p, 1] = 1 -> schedule[p, 2] != JAHO); 
constraint forall(p in P)(X[p, nb_days] = 1 -> schedule[p, nb_days-1] != JAHO); 
constraint forall(p in P, d in 2..nb_days-1)(X[p, d] = 1 -> schedule[p, d-1] != JAHO /\ schedule[p, d+1] != JAHO);           

% optimization objective: fairness
array[P] of var float: workload;
constraint forall(p in P)(workload[p] = sum(d in D)(X[p, d]) / (nb_days - card(F[p])));

var float: fairness_score;
constraint fairness_score = max(workload) - min(workload);

% optimization objective: balance
array[P, 1..nb_days] of var 0..nb_days: a;
constraint forall(p in P, d in D)(
             if d = 1
             then 
               if X[p, d] = 1
               then a[p, d] = 0
               else a[p, d] = min_balance
               endif
             else 
               if X[p, d] = 1 
               then a[p, d] = 0
               else a[p, d] = 1 + a[p, d-1]
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
       ["balance_score:"] ++ [show(min(z))] ++ ["\n"] ++
       ["fairness_score:" ++ show(fairness_score)++"\n"] ++ 
       [
         if d = 1 then 
             "assistant_id:" ++ show(personnel_id[p]) ++ " " 
          ++ "workload:" ++ show(workload[p]) ++ " " 
         else "" endif ++

         if schedule[p, d] == FREE 
         then if fix(X[p, d]) = 1 then "JAEV " else "FREE " endif 
         else "" 
         endif ++
         if schedule[p, d] == JANW then "JANW " else "" endif ++
         if schedule[p, d] == SAEW then "SAEW " else "" endif ++
         if schedule[p, d] == JAWE then "JAWE " else "" endif ++
         if schedule[p, d] == JAHO then "JAHO " else "" endif ++
         if schedule[p, d] == SAWE then "SAWE " else "" endif ++
         if schedule[p, d] == SAHO then "SAHO " else "" endif ++
         if schedule[p, d] == TPWE then "TPWE " else "" endif ++
         if schedule[p, d] == TPHO then "TPHO " else "" endif ++
         if schedule[p, d] == CALL then "CALL " else "" endif ++
         if d == nb_days then "\n" else "" endif 
         | p in P, d in D
       ];


