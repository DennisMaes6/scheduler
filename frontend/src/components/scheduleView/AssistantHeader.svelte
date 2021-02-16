<script lang=typescript>

    import type { Assistant } from '../../openapi'
    
    export let assistant: Assistant

</script>

<main>
    <div class="flex h-6 w-20 items-center justify-end">
        <p class="text-black text-xs font-bold"> {assistant.type} </p>
    </div>
</main>