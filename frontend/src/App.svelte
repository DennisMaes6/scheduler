<script lang=typescript>

    import ScheduleView from './components/schedule-view/ScheduleView.svelte';
    import ModelInput from './components/input/ModelInput.svelte';
    import { Service } from './openapi';

</script>

<main>
    <div class="flex flex-row">
        <div class="flex-none w-36 mt-10 ml-2">
            {#await Service.getModelParams()}
                <p> loading... </p>
            {:then modelParams} 
                <ModelInput {modelParams}/>
            {:catch error}
                <div>
                    <p style="color: red">! {error}: </p> 
                    <p style="color: red">{error.body} </p>
                </div>
            {/await}
        </div>
        <div class="flex-grow container">
            {#await Service.getSchedule()}
                <p>loading...</p>
            {:then schedule}
                <ScheduleView {schedule}/>
            {:catch error}
                <div>
                    <p style="color: red">! {error}: </p> 
                    <p style="color: red">{error.body} </p>
                </div>
            {/await}
        </div>
    </div>
</main>