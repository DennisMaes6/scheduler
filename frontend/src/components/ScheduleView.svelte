<script lang=typescript>

    import type { Schedule } from '../openapi';
    import { ShiftType } from '../openapi';

    export let schedule: Schedule;

    let days: number[] = Array.from({length: schedule.nb_days}, (_, i) => i + 1);

    function getColor(assignment: ShiftType): string {
        switch (assignment) {
            case ShiftType.JAEV:
                return "bg-yellow-400";
            case ShiftType.JANW:
                return "bg-green-400";
            case ShiftType.JAWH:
                return "bg-red-500";
            case ShiftType.SAEW:
                return "bg-green-500";
            case ShiftType.SAWH:
                return "bg-red-600";
            case ShiftType.CALL:
                return "bg-green-600";
            case ShiftType.TSPT:
                return "bg-red-700";
            case ShiftType.FREE:
                return "bg-white";
        }
    }

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

<!-- TODO: Make first column sticky -->
<main>
    <div class="flex flex-col overflow-scroll">
        <!-- Header with days of scheduling period -->
        <div class="sticky top-0 flex flex-row bg-white">
            <div class="flex-none w-32"></div>
            <div class="flex flex-col h-20 justify-center space-y-4">
                <div class="flex flex-row justify-center">
                    {#each days as day}
                        <div class="flex flex-none mx-1 w-16 items-center">
                            <p class="text-sm font-bold text-black"> {day % 7 == 1 ? "WEEK " + (Math.floor(day/7) + 1) : ""}</p>
                        </div>
                    {/each} 
                </div>
                <div class="flex flex-row items-center">
                    {#each days as day}
                        <div class="flex flex-none mx-1 w-16 items-center">
                            <p class="text-sm font-bold text-black"> {getDay(day)}</p>
                        </div>
                    {/each} 
                </div>
            </div>
        </div>
        <div class="flex flex-row">
            <!-- Assistant list -->
            <div class="flex flex-col w-24 mr-8">
                {#each schedule.assistants as assistant}
                    <div class="flex h-10 items-center justify-end">
                        <p class="text-black font-sm font-bold"> {assistant.type} </p>
                    </div>
                {/each}
            </div>
            <!-- Schedule -->
            <div class="flex flex-col">
                {#each schedule.individual_schedules as individual_schedule}
                    <div class="flex flex-row items-center">
                        {#each individual_schedule.assignments as assignment}
                            <div class="flex mx-1 my-1 w-16 h-8 {getColor(assignment)} rounded-lg justify-center items-center">
                                <p class="text-sm font-bold text-white"> {assignment}</p>
                            </div>
                        {/each}
                        <div class="bg-white w-20 text-white">
                            spacing
                        </div>
                    </div>
                {/each}
            </div>
        </div>
    </div>
</main>