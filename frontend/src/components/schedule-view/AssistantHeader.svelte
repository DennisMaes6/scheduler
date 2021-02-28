<script lang=typescript>

    import type { Assistant } from '../../openapi'
    
    export let assistant: Assistant

</script>

<main>
    <div class="flex flex-row h-6 items-center justify-end cursor-default">
        
        <p class="mx-2 w-12 text-gray-700 text-xs font-bold"> {assistant.workload} </p>
        <p class="mx-2 w-12 text-black text-xs font-bold"> {assistant.type} </p>
    </div>
</main>