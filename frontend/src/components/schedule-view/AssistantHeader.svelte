<script lang=typescript>

    import type { Assistant } from '../../openapi'
    import { AssistantType } from '../../openapi'
    
    export let assistant: Assistant
    export let workload: number
    export let max_workload: number
    export let min_workload: number

    function getWorkloadTextStyle() {
        if (workload == max_workload) {
            return "text-red-500 font-bold"
        } else if (workload == min_workload) {
            return "text-green-500 font-bold"
        } else {
            return "text-gray-700"
        }
    }
</script>


<div class="flex flex-row flex-none h-6 justify-between place-items-center cursor-default">
    {#if assistant.type === AssistantType.JA}
        <p class="text-xs mr-2">JA</p>
    {:else if assistant.type === AssistantType.JA_F}
        <p class="text-xs mr-2">JA<sub>F</sub></p>
    {:else if assistant.type === AssistantType.SA}
        <p class="text-xs mr-2">SA</p>
    {:else if assistant.type === AssistantType.SA_F}
        <p class="text-xs mr-2">SA<sub>F</sub></p>
    {:else if assistant.type === AssistantType.SA_NEO}
        <p class="text-xs mr-2">SA<sub>N</sub></p>
    {:else if assistant.type === AssistantType.SA_F_NEO}
        <p class="text-xs mr-2">SA<sub>FN</sub></p>
    {/if}
    <p class="{getWorkloadTextStyle()} text-xs"> {workload.toFixed(2)} </p>
    <p class="w-32 text-right text-black text-xs font-bold "> {assistant.name} </p>
</div>
