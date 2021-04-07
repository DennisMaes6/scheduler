<script lang=typescript>

    import ScheduleView from './components/schedule-view/ScheduleView.svelte';
    import ModelInput from './components/model-input/ModelInput.svelte';
    import { Service } from './openapi';
    import InstanceInput from './components/instance-input/InstanceInput.svelte';

</script>

<style>
.loader {
    border-top-color: #3498db;
    -webkit-animation: spinner 1.5s linear infinite;
    animation: spinner 1.5s linear infinite;
}

@-webkit-keyframes spinner {
    0% { -webkit-transform: rotate(0deg); }
    100% { -webkit-transform: rotate(360deg); }
}

@keyframes spinner {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
}
</style>

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
                <div class="mt-56">
                    <p class="mx-auto w-20"> Loading... </p>
                </div>
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