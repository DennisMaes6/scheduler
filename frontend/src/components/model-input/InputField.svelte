<script lang=typescript>

    export let value: number;
    export let step: number;

</script>


<main>
    <div class="rounded-md shadow-sm">
        <input type="number" bind:value={value} step={step} 
            class="pl-1 h-6 w-full rounded-md text-sm">
    </div>
</main>