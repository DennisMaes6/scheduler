<script lang=typescript>

    import ScheduleView from './components/scheduleView/ScheduleView.svelte';
    import ModelInput from './components/input/ModelInput.svelte';
    import { Service } from './openapi';

</script>

<main>
    <div class="flex flex-row">
        <div class="flex-none w-24 mt-10">
            <ModelInput/>
        </div>
        <div class="container">
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
    </div>
</main>