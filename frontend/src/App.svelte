<script lang=typescript>

    import ScheduleView from './components/schedule-view/ScheduleView.svelte';
    import ModelInput from './components/model-input/ModelInput.svelte';
    import { Service } from './openapi';
    import InstanceInput from './components/instance-input/InstanceInput.svelte';

</script>

<main>
    <div class="flex flex-row">
        <div class="flex flex-col space-y-5">
            <div class="flex-none w-36 ml-2">
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
            <div class="flex-none w-36 ml-2">
                {#await Service.getInstanceData()}
                    <p> loading... </p>
                {:then data} 
                    <InstanceInput {data}/>
                {:catch error}
                    <div>
                        <p style="color: red">! {error}: </p> 
                        <p style="color: red">{error.body} </p>
                    </div>
                {/await}
            </div>
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