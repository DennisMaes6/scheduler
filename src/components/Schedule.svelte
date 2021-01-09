<script lang=typescript>

    import type { ScheduleData } from '../types/schedule-data';

    export let schedule_data: ScheduleData;

</script>


<main>
    {#each schedule_data.individual_schedules as individual_schedule}
        <p> {individual_schedule.employee_nb} </p>
        <div>
        {#each individual_schedule.assignments as assignment}
            <div> {assignment} </div>
        {/each}
        </div>
    {/each}
</main>