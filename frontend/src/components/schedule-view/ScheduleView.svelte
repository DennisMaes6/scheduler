<script lang=typescript>

    import type { Schedule, InstanceData, Assignment, Assistant } from '../../openapi';
    
    import AssistantHeader from './AssistantHeader.svelte'
    import DayHeader from './DayHeader.svelte';
    import AssignmentBox from './AssignmentBox.svelte';

    export let schedule: Schedule;
    export let data: InstanceData;

    let max_workload = Math.max(...schedule.individual_schedules.map((a) => a.workload))
    let min_workload = Math.min(...schedule.individual_schedules.map((a) => a.workload))

    function getAssistant(id: number): Assistant {
        return data.assistants.find(a => a.id === id);
    }

    function get_assignment(assistant_id: number, day_id: number): Assignment {
        return schedule.individual_schedules
            .find((s) => s.assistant_id == assistant_id).assignments
            .find((a) => a.day_nb == day_id)
    }

</script>

<main class="ml-4">
    <div class="font-bold text-sm">
        <p> fairness score: {schedule.fairness_score.toFixed(2)} </p>
        <p> JAEV fairness score: {schedule.jaev_fairness_score.toFixed(2)} </p>
    </div>
    <div class="mt-5 flex flex-row">
        <!-- Assistant list -->
        <div class="flex flex-grow flex-col space-y-1 mr-4">
            <p class="mt-4 text-xs"> Workload </p>
            {#each schedule.individual_schedules as is}
                <AssistantHeader assistant={getAssistant(is.assistant_id)} workload={is.workload} {max_workload} {min_workload}/>
            {/each}
        </div>
        <div class="flex flex-col space-y-1 overflow-x-scroll">
            <!-- Header with days of scheduling period -->
            <div class="flex flex-row space-x-1">
                {#each data.days as day}
                    <DayHeader {day}/>
                    {#if day.id % 7 === 0}
                        <div class="flex flex-none w-4"/>
                    {/if}
                {/each}
            </div>
            <!-- Schedule -->
            {#each schedule.individual_schedules as is}
                <div class="flex flex-row space-x-1">
                    {#each data.days as day}
                        <AssignmentBox assignment={is.assignments.find(a => a.day_nb === day.id)}/>
                        {#if day.id % 7 === 0}
                            <div class="flex flex-none w-4"/>
                        {/if}
                    {/each}
                </div>
            {/each}
        </div>
    </div>
</main>