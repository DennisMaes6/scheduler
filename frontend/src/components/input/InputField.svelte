<script lang=typescript>

    export let label: string;
    export let value: number;
    export let step: number;

</script>


<main>
    <div>
        <label for="input" class="block text-sm font-medium text-gray-700">{label}</label>
        <div class="mt-1 relative rounded-md shadow-sm">
            <input type="number" value={value} step={step} 
                class="focus:ring-indigo-500 focus:border-indigo-500 block w-full pl-2 border-gray-300 rounded-md">
        </div>
    </div>
</main>