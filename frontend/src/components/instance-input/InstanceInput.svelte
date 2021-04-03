<script lang=typescript>
    
    import type { InstanceData } from "../../openapi";

    import { Service } from '../../openapi';

    import Modal from "../Modal.svelte";
    import Button from "../model-input/Button.svelte";
    import IconButton from "../model-input/IconButton.svelte";
    import InputField from "../model-input/InputField.svelte";
    import AssistantInput from "./AssistantInput.svelte";

    export let data: InstanceData

    function handleSubmit() {
        Service.postInstanceData(data)
    }
    
    function addHoliday() {
        data.holidays.push(0)
        data.holidays = data.holidays
    }

    function removeHoliday(i: number) {
        data.holidays.splice(i, 1)
        data.holidays = data.holidays
    }
</script>

<main>
    <form class="flex flex-col"> 
        <p class="font-semibold text-sm cursor-default"> Instance data </p>
        <p class="mt-1 font-semibold text-xs text-gray-500 cursor-default"> Number of weeks </p>
        <InputField bind:value={data.nb_weeks} />
        <p class="mt-1 font-semibold text-xs text-gray-500 cursor-default"> Holidays </p>
        <div class="flex flex-col space-y-1">
            {#each data.holidays as h, i (i)}
            <div class="flex flex-row space-x-2 justify-between">
                <div class="flex-grow">
                    <InputField bind:value={h} />
                </div>
                <IconButton callback={() => removeHoliday(i)}><span class="h-6 w-6 color-gray-500 iconify" data-icon="ic:round-remove-circle-outline" data-inline="false"></span></IconButton>
            </div>
        {/each}
        </div>
        <IconButton callback={addHoliday}><span class="h-6 w-6 color-gray-500 iconify" data-icon="ic:baseline-add-circle-outline" data-inline="false"></span></IconButton>
        <Modal>
            <div slot="trigger" let:open>
                <Button callback={open}> Edit assistants </Button>
            </div>
            <div class="my-2" slot="header">
                <h1 class="font-bold"> Edit assistants </h1>
            </div>
            <div class="my-2" slot="content">
                <AssistantInput bind:assistants={data.assistants} nb_weeks={data.nb_weeks}/>
            </div>
            <div class="my-5 flex flex-row justify-end space-x-2" slot="footer" let:close>
                <Button primary={false} callback={close}> Close </Button>
                <Button callback={() => {handleSubmit(); close()}}> Submit </Button>
            </div>
        </Modal>

        <div class="my-1 mx-auto">
            <Button callback={handleSubmit}> Submit </Button>
        </div>
    </form>
</main>