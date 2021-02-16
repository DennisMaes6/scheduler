
<script lang=typescript>

    export let id: string

</script>

<main>
<div class="relative w-10">
    <input type="checkbox" name="toggle" id={id} class="toggle-checkbox absolute w-6 h-6 rounded-full bg-white border-4 appearance-none cursor-pointer"/>
    <label for={id} class="toggle-label overflow-hidden h-6 rounded-full bg-gray-300 cursor-pointer"></label>
</div>
</main>