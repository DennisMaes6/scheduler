<script lang=typescript>

    import type { Assistant, AssistantType } from '../openapi'
    
    export let assistant: Assistant

</script>

<main>
    <div class="flex h-8 w-20 items-center justify-end">
        <p class="text-black font-sm font-bold"> {assistant.type} </p>
    </div>
</main>