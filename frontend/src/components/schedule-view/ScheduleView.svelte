<script lang=typescript>

    import type { Schedule } from '../../openapi';
    
    import AssistantHeader from './AssistantHeader.svelte'
    import Assignment from './Assignment.svelte';
    import DayHeader from './DayHeader.svelte';

    export let schedule: Schedule;

    let days: number[] = Array.from({length: schedule.nb_days}, (_, i) => i + 1)

    let max_workload = Math.max(...schedule.assistants.map((a) => a.workload))
    let min_workload = Math.min(...schedule.assistants.map((a) => a.workload))

</script>

<main>
    <div class="flex flex-row space-x-10 ml-12 my-5 font-bold text-sm">
        <div>
            <p> fairness score: {schedule.fairness_score.toFixed(2)} </p>
            <p> balance score: {schedule.balance_score} </p>
        </div>
        <div>
            <p> JAEV fairness score: {schedule.jaev_fairness_score.toFixed(2)} </p>
            <p> JAEV balance score: {schedule.jaev_balance} </p>
        </div>
    </div>
    <div class="m-10 pl-6 flex flex-row pr-2">
        <!-- Assistant list -->
        <div class="flex flex-col w-32 mx-4 space-y-2">
            <p class="mx-2 mt-4 text-xs"> Load </p>
            {#each schedule.assistants as assistant}
                <AssistantHeader {assistant} max_workload={assistant.workload === max_workload} min_workload={assistant.workload === min_workload}/>
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