<script lang=typescript>

    import ScheduleView from './components/ScheduleView.svelte';

    import type { Schedule } from './openapi';
    import { Service } from './openapi';
    
    let schedule: Schedule;
    let loaded: boolean = false;
    

</script>

<main>
    {#await Service.getService()}
	    <p>loading...</p>
    {:then schedule}
	    <ScheduleView {schedule}/>
    {:catch error}
	    <p style="color: red">! {error.message}: </p> 
        <p style="color: red">{error.body} </p>
    {/await}
</main>