<script lang=typescript>

    import type { ScheduleData } from '../types/schedule-data';
    import { ShiftType } from '../types/schedule-data';


    export let schedule_data: ScheduleData;

    let days: number[] = Array.from({length: schedule_data.nb_days}, (_, i) => i + 1);

    function getColor(assignment: ShiftType): string {
        switch (assignment) {
            case ShiftType.JA_E:
                return "bg-yellow-400";
            case ShiftType.JA_NW:
                return "bg-green-400";
            case ShiftType.JA_WH:
                return "bg-red-500";
            case ShiftType.SA_EW:
                return "bg-green-500";
            case ShiftType.SA_WH:
                return "bg-red-600";
            case ShiftType.C:
                return "bg-green-600";
            case ShiftType.TS:
                return "bg-red-700";
            case ShiftType.O:
                return "bg-white";
        }
    }
</script>

<!-- TODO: Make first column sticky -->
<main>
    <div>
        <div class="flex-col">
            <div class="mb-2 sticky top-0 flex flex-row w-screen space-x-5 items-center bg-white">
                <div class="flex flex-none w-24 h-8 bg-white"></div>
                {#each days as day}
                    <div class="flex flex-none w-16 h-8 justify-center place-items-center bg-white">
                        <p class="text-sm font-bold text-black"> day {day}</p>
                    </div>
                {/each} 
            </div>
            {#each schedule_data.individual_schedules as individual_schedule}
            <div class="mb-3 flex flex-row flex-none space-x-5 items-center">
                <div class="sticky flex flex-none w-24 h-8 justify-center place-items-center">
                    <p class="text-sm font-bold text-black"> assistant { individual_schedule.assistant_id }</p>
                </div>
                {#each individual_schedule.assignments as assignment}
                    <div class="flex flex-none w-16 h-8 {getColor(assignment)} rounded-lg justify-center place-items-center">
                        <p class="text-sm font-bold text-white"> {assignment}</p>
                    </div>
                {/each}
                <div class="bg-white w-20 text-white">
                    spacing
                </div>
            </div>
        {/each}
        </div>
        <div class="bg-white h-20"></div>
    </div>
</main>