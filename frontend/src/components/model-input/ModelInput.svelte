<script lang=typescript>
    
    import type { ModelParameters} from "../../openapi";
    
    import { Service } from '../../openapi';

    import Assignment from "../schedule-view/Assignment.svelte";
    import Button from "./Button.svelte";
    import InputField from "./InputField.svelte";
    import Toggle from "./Toggle.svelte";

    export let modelParams: ModelParameters

    function handleSubmit() {
        Service.postModelParams(modelParams)
    }

</script>

<main>
    <form class="flex flex-col"> 
        <p class="font-semibold text-sm cursor-default"> Model parameters </p>
        <p class="mt-2 font-semibold text-xs text-gray-500 cursor-default"> Minimun balance score </p>
        <InputField bind:value={modelParams.balance_minimum} step={1} />

        <p class="mt-2 font-semibold text-xs text-gray-500 cursor-default"> fairness weight + max buffer </p>
        <div class="mt-2 flex flex-col space-y-2">
            {#each modelParams.shift_type_params as stp}
                <div class="flex flex-row space-x-1 items-center">
                    <Assignment assignment={{shift_type:stp.shift_type, part_of_min_balance:false}}/>
                    <InputField bind:value={stp.shift_workload} step={0.01}/>
                    <InputField bind:value={stp.max_buffer} step={1}/>
                </div>
            {/each}
        </div>

        <div class="my-5 mx-auto">
            <Button callback={handleSubmit}> Submit </Button>
        </div>
    </form>

</main>