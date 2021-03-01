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
  JAWH,
  SAEW,
  SAWH,
  TSPT,
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
constraint forall(d in D)(sum(p in P)(X[p, d]) = R[1 + ((d - 1) mod 7)]);

% HC3 - no assignments on free days
constraint forall(p in P, d in D)(X[p, d] = 1 -> not (d in F[p]));

% HC4 - no consecutive jaev assignements
constraint forall(p in P, d in 2..nb_days)(X[p, d-1] = 1 -> X[p, d] = 0); 

% HC5 - no assignments in the weeks before or after night weeks
constraint forall(p in P, d in 1..7)(X[p, d] = 1 -> schedule[p, d+7] != JANW); 
constraint forall(p in P, d in nb_days-6..nb_days)(X[p, d] = 1 -> schedule[p, d-7] != JANW); 
constraint forall(p in P, d in 8..nb_days-7)(X[p, d] = 1 -> schedule[p, d-7] != JANW /\ schedule[p, d+7] != JANW); 
           

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
       ["min_balance:"] ++ [show(min_balance)] ++ ["\n"] ++
       ["fairness_score:" ++ show(fairness_score)++"\n"] ++ 
       ["nb_days:"] ++ [show(nb_days)] ++ ["\n"] ++
       [
         if d = 1 then 
             "assistant:" ++ show(p) ++ " " 
          ++ "type:JA " 
          ++ "workload:" ++ show(workload[p]) ++ " " 
         else "" endif ++

         if schedule[p, d] == FREE 
         then if fix(X[p, d]) = 1 then "JAEV " else "FREE " endif 
         else "" 
         endif ++
         if schedule[p, d] == JANW then "JANW " else "" endif ++
         if schedule[p, d] == SAEW then "SAEW " else "" endif ++
         if schedule[p, d] == JAWH then "JAWH " else "" endif ++
         if schedule[p, d] == SAWH then "SAWH " else "" endif ++
         if schedule[p, d] == TSPT then "TSPT " else "" endif ++
         if schedule[p, d] == CALL then "CALL " else "" endif ++
         if d == nb_days then "\n" else "" endif 
         | p in P, d in D
       ];


