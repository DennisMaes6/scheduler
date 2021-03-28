<script lang=typescript>

    export let day: number
    export let holiday: boolean

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
    <div class="flex flex-col h-8 w-12 justify-end items-center cursor-default">
        {#if day % 7 == 1}
            <p class="text-xs font-bold text-black"> WEEK {(Math.floor(day/7) + 1)}</p>
        {/if}
        <div class="flex flex-row justify-items-center w-12 rounded {holiday ? "bg-yellow-200" : ""}">
            <p class="m-auto text-xs font-bold text-black"> {getDay(day)}</p>
        </div>
    </div>
</main>