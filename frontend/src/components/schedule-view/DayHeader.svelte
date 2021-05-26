<script lang=typescript>

    import type { Day } from '../../openapi';
    
    export let day: Day

    function day_of_week(day: Day): string {
        switch (day.id % 7) {
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
    <div class="flex flex-col flex-none h-8 w-12 justify-end items-center cursor-default">
        <div class="flex w-12 rounded">
            <p class="m-auto text-xs font-bold text-black"> {day_of_week(day)}</p>
        </div>
        <div class="flex w-12 rounded {day.is_holiday ? "bg-yellow-200" : ""}">
            <p class="m-auto text-xs font-bold text-black"> {day.date.day + "/" + day.date.month}</p>
        </div>
    </div>
</main>