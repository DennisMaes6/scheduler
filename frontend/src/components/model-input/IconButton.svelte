<script lang=typescript>

    export let callback: () => void
    
</script>

<button on:click={callback} type="button" class="focus:outline-none border-transparent">
    <slot/>
</button>