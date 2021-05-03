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

<main>
    <div class="ml-12 font-bold text-sm">
        <p> fairness score: {schedule.fairness_score.toFixed(2)} </p>
        <p> JAEV fairness score: {schedule.jaev_fairness_score.toFixed(2)} </p>
    </div>
    <div class="ml-10 mt-5 flex flex-row pr-2">
        <!-- Assistant list -->
        <div class="flex flex-col w-32 mx-4 space-y-2">
            <p class="mt-4 text-xs"> Workload </p>
            {#each schedule.individual_schedules as is}
                <AssistantHeader assistant={getAssistant(is.assistant_id)} workload={is.workload} {max_workload} {min_workload}/>
            {/each}
        </div>
        <div class="flex flex-col space-y-2 overflow-x-scroll">
            <!-- Header with days of scheduling period -->
            <div class="flex flex-row space-x-2">
                {#each data.days as day}
                    <DayHeader {day}/>
                {/each}
            </div>
            <!-- Schedule -->
            {#each schedule.individual_schedules as is}
                <div class="flex flex-row space-x-2">
                    {#each data.days as day}
                        <AssignmentBox assignment={is.assignments.find(a => a.day_nb === day.id)}/>
                    {/each}
                </div>
            {/each}
        </div>
    </div>
</main>