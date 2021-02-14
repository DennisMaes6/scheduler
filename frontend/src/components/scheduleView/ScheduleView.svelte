<script lang=typescript>

    import type { Schedule } from '../../openapi';
    
    import AssistantHeader from './AssistantHeader.svelte'
    import Assignment from './Assignment.svelte';
    import DayHeader from './DayHeader.svelte';

    export let schedule: Schedule;

    let days: number[] = Array.from({length: schedule.nb_days}, (_, i) => i + 1)

</script>

<main>
    <div class="flex flex-row">
        <!-- Assistant list -->
        <div class="flex flex-col mx-4 space-y-2">
            <div class="h-12"></div> <!-- spacing -->
            {#each schedule.assistants as assistant}
                <AssistantHeader {assistant}/>
            {/each}
        </div>
        <div class="flex flex-col space-y-2 overflow-x-scroll">
            <!-- Header with days of scheduling period -->
            <div class="flex flex-row space-x-2">
                {#each days as day}
                    <DayHeader {day}/>
                {/each}
            </div>
            <!-- Schedule -->
            {#each schedule.individual_schedules as individual_schedule}
                <div class="flex flex-row space-x-2">
                    {#each individual_schedule.assignments as assignment}
                        <Assignment {assignment}/>
                    {/each}
                </div>
            {/each}
        </div>
    </div>
</main>