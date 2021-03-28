
<script lang=typescript>

    export let callback: () => void
    
</script>

<button on:click={callback} type="button" class="bg-green-500 hover:bg-green-700 text-white font-bold py-1 px-2 rounded text-xs focus:outline-none border-transparent">
    <slot/>
</button>