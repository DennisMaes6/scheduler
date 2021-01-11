<script lang=typescript>

    import type { ScheduleData } from './types/schedule-data';
    import Schedule from './components/Schedule.svelte';
    
    let schedule_data: ScheduleData;
    let loaded: boolean = false;

    fetch("http://localhost:8080/schedule")
        .then(response => response.json())
        .then(result => {
            console.log(result)
            return result
        })
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