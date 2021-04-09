
<script lang=typescript>

    import type { InstanceData, Assistant } from "../../openapi";
    import { AssistantType } from "../../openapi";

    import IconButton from "../model-input/IconButton.svelte";
    import Button from "../model-input/Button.svelte";
    import Modal from "../Modal.svelte";
    import DayHeader from "../schedule-view/DayHeader.svelte";
    import NewAssistant from "./NewAssistant.svelte";

    export let data: InstanceData;

    const types: AssistantType[] = Object.values(AssistantType)

    let new_assistant: Assistant = {
        id: Math.max(...data.assistants.map((a) => a.id)) + 1,
        name: "",
        type: AssistantType.JA,
        free_days: []
    }

    function removeAssistant(assistant: Assistant) {
        data.assistants = data.assistants.filter((a) => a.id != assistant.id)
    }

    function addAssistant() {
        data.assistants.push(new_assistant)
        new_assistant = {
            id: Math.max(...data.assistants.map((a) => a.id)) + 1,
            name: "",
            type: AssistantType.JA,
            free_days: []
        }

        data.assistants = data.assistants
    }

    function toggle(assistant: Assistant, day_nb: number) {
        if (assistant.free_days.includes(day_nb)) {
            assistant.free_days = assistant.free_days.filter((d: number) => d != day_nb)
        } else {
            assistant.free_days.push(day_nb)
        }
        data.assistants = data.assistants
    }

</script>


<div class="flex flex-col space-y-1">
    <div class="flex flex-row">
        <div class="flex flex-none w-48">
            <Modal>
                <div slot="trigger" let:open>
                    <Button callback={open}> New </Button>
                </div>
                <div class="my-2" slot="header">
                    <h1 class="font-bold"> New assistant </h1>
                </div>
                <div class="my-2" slot="content">
                    <NewAssistant bind:assistant={new_assistant}/>
                </div>
                <div class="my-5 flex flex-row justify-end space-x-2" slot="footer" let:close>
                    <Button primary={false} callback={close}> Close </Button>
                    <Button callback={() => {addAssistant(); close()}}> Create </Button>
                </div>
            </Modal>
        </div>
        <div class="flex flex-none w-11"></div>
        <div class="flex flex-row space-x-1">
            {#each data.days as day}
                <DayHeader {day}/>
            {/each}
        </div>

    </div>
    {#each types as type}
        {#each data.assistants.filter((a) => a.type == type) as assistant (assistant.id)}
            <div class="flex flex-row space-x-1">
                <IconButton callback={() => removeAssistant(assistant)}>
                    <span class="h-4 w-X color-gray-500 iconify" 
                          data-icon="ic:round-remove-circle-outline" 
                          data-inline="false"/>
                </IconButton>
                <p class="flex flex-none w-32 text-black text-sm"> {assistant.name} </p>
                <p class="flex flex-none w-20 text-black text-sm"> {assistant.type} </p>
                {#each data.days as day}
                    <div on:click={() => toggle(assistant, day.id)} class="flex-none w-12 h-6 {assistant.free_days.includes(day.id) ? "bg-yellow-400" : "bg-gray-100"} rounded cursor-pointer flex justify-center items-center">
                        <div class="text-xs text-white font-bold"> FREE </div>
                    </div>
                {/each}
            </div>
        {/each}
    {/each}
</div>