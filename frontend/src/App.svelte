<script lang=typescript>

    import ScheduleView from './components/schedule-view/ScheduleView.svelte';
    import ModelInput from './components/model-input/ModelInput.svelte';
    import { Service } from './openapi';
    import InstanceInput from './components/instance-input/InstanceInput.svelte';

</script>

<div class="flex flex-row flex-nowrap h-full w-full space-x-5">
    <div class="flex flex-none flex-col h-full overflow-y-scroll">
        <div class="w-36 ml-2">
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
        <div class="w-36 ml-2">
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

    <div class="flex flex-col h-full overflow-scroll space-y-5">
        {#await Promise.all([Service.getFileSchedule(), Service.getInstanceData()])}
            <div class="flex h-full w-full">
                <p class="m-auto w-20"> Loading... </p>
            </div>
        {:then result}
            <ScheduleView schedule={result[0]} data={result[1]}/>
        {:catch error}
            <div>
                <p style="color: red">! {error}: </p> 
                <p style="color: red">{error.body} </p>
            </div>
        {/await}
    </div>
</div>