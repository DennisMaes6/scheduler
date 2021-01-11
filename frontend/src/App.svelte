<script lang=typescript>

    import type { ScheduleData } from './types/schedule-data';
    import Schedule from './components/Schedule.svelte';
    
    let schedule_data: ScheduleData;
    let loaded: boolean = false;

    fetch("data.json")
        .then(response => response.json())
        .then(result => {
            schedule_data = result;
            loaded = true; 
        });

</script>

<main>
    {#if !loaded}
        <p> loading... </p>
    {:else}
        <Schedule {schedule_data}/>
    {/if}
</main>