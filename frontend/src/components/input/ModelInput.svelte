<script lang=typescript>
    
    import { loop_guard, stop_propagation } from "svelte/internal";
import type { ModelParameters } from "../../openapi";
    import { ShiftType } from "../../openapi";
    
    import Assignment from "../scheduleView/Assignment.svelte";
import Button from "./Button.svelte";
    import InputField from "./InputField.svelte";
    import Toggle from "./Toggle.svelte";

    // mock data
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
                shift_type: ShiftType.JAWH,
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
                included_in_balance: false,
                fairness_weight: 0,
            },
        ]
    }

    function handlePrint() {
        console.log(modelParams)
    }

</script>

<main>
    <form class="flex flex-col"> 
        <p class="font-semibold text-sm cursor-default"> Minimun balance score: </p>
        <InputField value={modelParams.balance_minimum} step={1} />

        <p class="mt-12 font-semibold text-sm cursor-default"> Shift type specific paramaters: </p>
        <p class="mt-2 font-semibold text-xs text-gray-500 cursor-default"> fairness weight + included in balance </p>
        <div class="mt-2 flex flex-col space-y-2">
            {#each modelParams.shift_type_params as stp}
                <div class="flex flex-row space-x-1 items-center">
                    <Assignment assignment={stp.shift_type}/>
                    <InputField bind:value={stp.fairness_weight} step={0.01}/>
                    <Toggle bind:checked={stp.included_in_balance} id={stp.shift_type}/>
                </div>
            {/each}
        </div>

        <div class="my-5 mx-auto">
            <Button callback={handlePrint}> Submit </Button>
        </div>
    </form>

</main>