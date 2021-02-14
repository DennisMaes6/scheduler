<script lang=typescript>
    
    import { ModelParameters, ShiftType } from "../../openapi";

    let modelParams: ModelParameters = {
        balance_minimum: 3,
        balance_shifts: [ShiftType.JAEV, ShiftType.JANW, ShiftType.TSPT],
        fairness_shift_weights: [
            {
                shift_type: ShiftType.JAEV,
                weight: 0,
            },
            {
                shift_type: ShiftType.JANW,
                weight: 1,
            },
            {
                shift_type: ShiftType.TSPT,
                weight: 2
            }
        ]

        
    }

</script>

<main>
    
    <p> test </p>
</main>