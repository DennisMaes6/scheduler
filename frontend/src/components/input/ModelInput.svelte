<script lang=typescript>
    
    import { ModelParameters, ShiftType } from "../../openapi";

    let modelParams: ModelParameters = {
        balance_minimum: 3,
        shift_type_params: [
            {
                shift_type: ShiftType.JAEV,
                included_in_balance: true,
                fairness_weight: 0,
            },
            {
                shift_type: ShiftType.JANW,
                included_in_balance: true,
                fairness_weight: 0,
            },
            {
                shift_type: ShiftType.JANW,
                included_in_balance: true,
                fairness_weight: 0,
            },
            {
                shift_type: ShiftType.SAEW,
                included_in_balance: true,
                fairness_weight: 0,
            },
            {
                shift_type: ShiftType.SAWH,
                included_in_balance: true,
                fairness_weight: 0,
            },
            {
                shift_type: ShiftType.CALL,
                included_in_balance: true,
                fairness_weight: 0,
            },
            {
                shift_type: ShiftType.TSPT,
                included_in_balance: true,
                fairness_weight: 0,
            },
        ]
    }

</script>

<main>
    
    <p> test </p>
</main>