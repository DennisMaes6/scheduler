<script lang=typescript>
    
    import type { InstanceData, AssistantInstance } from "../../openapi";

    import { AssistantType } from "../../openapi";
    import { Service } from '../../openapi';

    import Modal from "../Modal.svelte";
    import Button from "../model-input/Button.svelte";
    import IconButton from "../model-input/IconButton.svelte";
    import InputField from "../model-input/InputField.svelte";

    export let data: InstanceData

    const types: AssistantType[] = [
        AssistantType.JA,
        AssistantType.JA_F,
        AssistantType.SA,
        AssistantType.SA_F,
        AssistantType.SA_NEO,
        AssistantType.SA_F_NEO
    ]

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

    function removeAssistant(assistant: AssistantInstance) {
        data.assistants = data.assistants.filter((a) => a.id != assistant.id)
        let index = 1
        data.assistants.forEach((a) => {
            a.id = index
            index++
        })
    }

    function addAssistant(type: AssistantType) {
        data.assistants.push(
            {
                id: Math.max(...data.assistants.map((a) => a.id)) + 1,
                type,
                free_days: []
            }
        )
        let index = 1
        data.assistants.forEach((a) => {
            a.id = index
            index++
        })
        data.assistants = data.assistants
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
                <div class="flex flex-col space-y-2">
                    {#each types as type}
                        <h1 class="my-2"> {type} </h1>
                        {#each data.assistants.filter((a) => a.type == type) as assistant (assistant.id)}
                            <div class="flex flex-row space-x-2">
                                <p class="w-16 text-black text-sm"> {assistant.type} </p>
                                <Modal>
                                    <div slot="trigger" let:open>
                                        <Button callback={open}> Edit </Button>
                                    </div>
                                    <div class="my-2" slot="content">
                                        test
                                    </div>
                                </Modal>
                                
                                <IconButton callback={() => removeAssistant(assistant)}>
                                    <span class="h-6 w-6 color-gray-500 iconify" 
                                          data-icon="ic:round-remove-circle-outline" 
                                          data-inline="false"/>
                                </IconButton>
                            </div>
                        {/each}
                        <Button callback={() => addAssistant(type)}> New </Button>
                    {/each}
  
                </div>
            </div>
            <div class="my-5 flex flex-row justify-end space-x-2" slot="footer" let:close>
                <Button primary={false} callback={close}> Close </Button>
                <Button callback={handleSubmit}> Submit </Button>
            </div>

        </Modal>
        <div class="my-1 mx-auto">
            <Button callback={handleSubmit}> Submit </Button>
        </div>
    </form>
</main>