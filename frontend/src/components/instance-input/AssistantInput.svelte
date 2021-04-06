
<script lang=typescript>

    import type { AssistantInstance } from "../../openapi";
    import { AssistantType } from "../../openapi";

    import IconButton from "../model-input/IconButton.svelte";
    import Button from "../model-input/Button.svelte";
    import Modal from "../Modal.svelte";
    import DayHeader from "../schedule-view/DayHeader.svelte";
    import NewAssistant from "./NewAssistant.svelte";

    export let assistants: AssistantInstance[]
    export let nb_weeks: number;

    let days: number[] = Array.from({length: nb_weeks*7}, (_, i) => i + 1)

    const types: AssistantType[] = Object.values(AssistantType)

    let new_assistant: AssistantInstance = {
        id: assistants.length + 1,
        name: "",
        type: AssistantType.JA,
        free_days: []
    }

    function removeAssistant(assistant: AssistantInstance) {
        assistants = assistants.filter((a) => a.id != assistant.id)
        assistants.forEach((a, i) => {
            a.id = i+1
        })
    }

    function addAssistant() {
        assistants.push(new_assistant)
        new_assistant = {
            id: assistants.length + 1,
            name: "",
            type: AssistantType.JA,
            free_days: []
        }
        assistants.forEach((a, i) => {
            a.id = i+1
        })
        assistants.sort((a, b) => {
            if (a.type < b.type) return -1;
            if (a.type > b.type) return 1;
            return 0;
        });
        assistants = assistants
    }

    function toggle(assistant: AssistantInstance, day: number) {
        console.log(day)
        console.log(assistant.free_days)
        if (assistant.free_days.includes(day)) {
            assistant.free_days = assistant.free_days.filter((d: number) => d != day)
        } else {
            assistant.free_days.push(day)
        }
        console.log(assistant.free_days)
        assistants.find((a) => a.id == assistant.id).free_days = assistant.free_days
        assistants = assistants
    }
</script>


<div class="flex flex-col space-y-2">
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
        <div class="flex flex-none w-14"></div>
        {#each days as day}
            <DayHeader {day}/>
        {/each}
    </div>
    {#each types as type}
        {#each assistants.filter((a) => a.type == type) as assistant (assistant.id)}
            <div class="flex flex-row space-x-2">
                <IconButton callback={() => removeAssistant(assistant)}>
                    <span class="h-6 w-6 color-gray-500 iconify" 
                          data-icon="ic:round-remove-circle-outline" 
                          data-inline="false"/>
                </IconButton>
                <p class="flex flex-none w-32 text-black text-sm"> {assistant.name} </p>
                <p class="flex flex-none w-20 text-black text-sm"> {assistant.type} </p>
                {#each days as day}
                    {#if assistant.free_days.includes(day)}
                        <div on:click={() => toggle(assistant, day)} class="flex flex-none w-10 h-6 bg-yellow-400 rounded-lg justify-center items-center cursor-default">
                            <div class="text-xs text-white font-bold"> FREE </div>
                        </div>
                    {:else}
                        <div on:click={() => toggle(assistant, day)} class="flex flex-none w-10 h-6 bg-gray-200 rounded-lg justify-center items-center cursor-default">
                        </div>
                    {/if}
                {/each}
            </div>
        {/each}
    {/each}
</div>