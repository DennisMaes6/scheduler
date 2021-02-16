<script lang=typescript>

    export let day: number

    function getDay(dayNb: number): string {
        switch (dayNb % 7) {
            case 0:
                return "Thu";
            case 1:
                return "Fri";
            case 2:
                return "Sat";
            case 3:
                return "Sun";
            case 4:
                return "Mon";
            case 5:
                return "Tue";
            case 6:
                return "Wed";
        }
    }

</script>


<main>
    <div class="flex flex-col h-12 w-12 justify-end">
        {#if day % 7 == 1}
            <p class="text-xs font-bold text-black"> WEEK {(Math.floor(day/7) + 1)}</p>
        {/if}
        <p class="text-xs font-bold text-black"> {getDay(day)}</p>
    </div>
</main>