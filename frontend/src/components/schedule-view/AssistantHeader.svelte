<script lang=typescript>

    import type { Assistant } from '../../openapi'
    
    export let assistant: Assistant

</script>

<main>
    <div class="flex flex-row h-6 items-center justify-end cursor-default">
        <div class="flex flex-row w-20">
            {#each assistant.highest_unfair_workload as huw}
                <p class="mx-1 text-gray-700 text-xs font-bold"> {huw} </p>
            {/each}
        </div>

        <div class="flex flex-row w-20">
            {#each assistant.lowest_unfair_workload as luw}
                <p class="mx-1 text-gray-700 text-xs font-bold"> {luw} </p>
            {/each}
        </div>

        
        <p class="w-12 text-black text-xs font-bold"> {assistant.type} </p>
    </div>
</main>