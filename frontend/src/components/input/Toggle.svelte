
<script lang=typescript>

    export let id: string
    export let checked: boolean

</script>

<main>
<div class="relative w-10 select-none">
    <input bind:checked={checked} type="checkbox" name="toggle" id={id} class="toggle-checkbox absolute block w-6 h-6 rounded-full bg-white border-4 appearance-none cursor-pointer focus:outline-none"/>
    <label for={id} class="toggle-label overflow-hidden h-6 rounded-full bg-gray-300 cursor-pointer"></label>
</div>
</main>