<script lang=typescript>

    import type { Assistant } from '../../openapi'
    
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

<main>
    <div class="flex flex-row h-6 justify-between place-items-center cursor-default">
        <p class="mx-2 {getWorkloadTextStyle()} text-xs"> {workload.toFixed(2)} </p>
        <p class="mx-2 text-black text-xs font-bold"> {assistant.type} </p>
    </div>
</main>