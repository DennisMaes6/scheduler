<script lang=typescript>

    import type { Schedule } from '../../openapi';
    
    import AssistantHeader from './AssistantHeader.svelte'
    import Assignment from './Assignment.svelte';
    import DayHeader from './DayHeader.svelte';

    export let schedule: Schedule;

    let days: number[] = Array.from({length: schedule.nb_days}, (_, i) => i + 1)

</script>

<main>
    <div class="ml-6 my-5 font-bold text-sm">
        <p> fairness score: {schedule.fairness_score} </p>
        <p> balance score: {schedule.balance_score} </p>
    </div>
    <div class="flex flex-row pr-2">
        <!-- Assistant list -->
        <div class="flex flex-col mx-4 space-y-2">
            <div class="flex flex-row h-12">
                <p class="mx-2 w-16 text-xs"> High workload </p>
                <p class="mx-2 w-16 text-xs"> Low workload </p>
            </div> 
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