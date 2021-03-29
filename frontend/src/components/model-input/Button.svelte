
<script lang=typescript>

    export let callback: () => void
    export let primary: boolean = true
    
</script>

{#if primary}
    <button on:click={callback} type="button" class="bg-green-500 hover:bg-green-700 text-white font-bold py-1 px-2 rounded text-xs focus:outline-none border-transparent">
        <slot/>
    </button>
{:else}
    <button on:click={callback} type="button" class="hover:bg-gray-100 text-green-500 font-bold py-1 px-2 rounded text-xs focus:outline-none border-transparent">
        <slot/>
    </button>
{/if}