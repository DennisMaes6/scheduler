<script lang=typescript>

    import ScheduleView from './components/scheduleView/ScheduleView.svelte';
    import ModelInput from './components/input/ModelInput.svelte';
    import { Service } from './openapi';

</script>

<main>
    <div class="flex flex-row">
        <ModelInput/>
        {#await Service.getService()}
            <p>loading...</p>
        {:then schedule}
            <ScheduleView {schedule}/>
        {:catch error}
            <div>
                <p style="color: red">! {error.message}: </p> 
                <p style="color: red">{error.body} </p>
            </div>
        {/await}
    </div>
</main>