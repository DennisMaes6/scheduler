
<script lang=typescript>

    import type { InstanceData, Assistant } from "../../openapi";
    import { AssistantType } from "../../openapi";

    import IconButton from "../model-input/IconButton.svelte";
    import Button from "../model-input/Button.svelte";
    import Modal from "../Modal.svelte";
    import DayHeader from "../schedule-view/DayHeader.svelte";
    import NewAssistant from "./NewAssistant.svelte";

    import { onMount } from 'svelte';

    onMount(async () => {
		let isSyncingLeftScroll: boolean = false;
        let isSyncingTopScroll: boolean = false;
        let isSyncingCenterScroll: boolean = false;
        let leftDiv: HTMLElement = document.getElementById('left');
        let topDiv: HTMLElement = document.getElementById('top');
        let centerDiv: HTMLElement = document.getElementById('center');

        leftDiv.onscroll = function() {
            if (!isSyncingLeftScroll) {
                isSyncingCenterScroll = true;
                centerDiv.scrollTop = leftDiv.scrollTop;
            }
            isSyncingLeftScroll = false;
        }

        topDiv.onscroll = function() {
            if (!isSyncingTopScroll) {
                isSyncingCenterScroll = true;
                centerDiv.scrollLeft = topDiv.scrollLeft;
            }
            isSyncingTopScroll = false;
        }

        centerDiv.onscroll = function() {
            if (!isSyncingCenterScroll) {
                isSyncingLeftScroll = true;
                isSyncingTopScroll = true;
                leftDiv.scrollTop = centerDiv.scrollTop;
                topDiv.scrollLeft = centerDiv.scrollLeft;
            }
            isSyncingCenterScroll = false;

        }
	});

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

    function setWeek(assistant: Assistant, weekNb: number, selected: boolean) {
        if (selected) {
            [7*weekNb+1, 7*weekNb+2, 7*weekNb+3, 7*weekNb+4, 7*weekNb+5, 7*weekNb+6, 7*weekNb+7]
                .forEach((day) => {
                    if (!assistant.free_days.includes(day)) assistant.free_days.push(day)
                })
        } else {
            [7*weekNb+1, 7*weekNb+2, 7*weekNb+3, 7*weekNb+4, 7*weekNb+5, 7*weekNb+6, 7*weekNb+7]
                .forEach((day) => {
                    assistant.free_days = assistant.free_days.filter((d: number) => d != day) 
                })   
        }
        data.assistants = data.assistants 
    }

    function freeWeek(assistant: Assistant, weekNb: number) {
        return [7*weekNb+1, 7*weekNb+2, 7*weekNb+3, 7*weekNb+4, 7*weekNb+5, 7*weekNb+6, 7*weekNb+7]
            .every((day) => assistant.free_days.includes(day))
    }
</script>


<div class="grid grid-cols-12" style="height: 70vh; width: 80vw">
    <div class="flex-none mb-2 col-span-3">
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
    <div id="top" class="scrollbar-hidden flex flex-row col-span-9 space-x-1 overflow-auto">
        {#each data.days as day}
            {#if day.id % 7 === 1}
                <!-- placeholder -->
                <div class="flex flex-none w-8"/>
            {/if}
            <DayHeader {day}/>
        {/each}
    </div>
    <div id="left" class="scrollbar-hidden flex flex-col space-y-1 col-span-3 overflow-auto">
        {#each types as type}
            {#each data.assistants.filter((a) => a.type == type) as assistant (assistant.id)}
                <div class="p-2 flex-none flex h-6 flex-row w-full justify-between">
                    <IconButton callback={() => removeAssistant(assistant)}>
                        <span class="h-4 w-8 color-gray-500 iconify" 
                            data-icon="ic:round-remove-circle-outline" 
                            data-inline="false"/>
                    </IconButton>
                    <p class="w-20 text-black text-xs"> {assistant.type} </p>
                    <p class="font-bold w-32 h-6 text-black text-xs text-right truncate"> {assistant.name} </p>
                </div>   
            {/each}
        {/each}
    </div>
    <div id="center" class="scrollable flex flex-col col-span-9 space-y-1 overflow-auto">
        {#each types as type}
            {#each data.assistants.filter((a) => a.type == type) as assistant (assistant.id)}
                <div class="flex flex-row space-x-1">
                    {#each data.days as day}
                        {#if day.id % 7 === 1}
                        <div class="flex w-8 justify-items-end items-center">
                            <input type="checkbox" checked={freeWeek(assistant, (day.id-1)/7)} class="ml-4 w-4 text-xs text-black font-bold" on:change={e => setWeek(assistant, (day.id-1)/7, e.target.checked)}/>
                        </div>
                        {/if}
                        <div on:click={() => toggle(assistant, day.id)} class="flex-none w-12 h-6 {assistant.free_days.includes(day.id) ? "bg-yellow-400" : "bg-gray-100"} rounded cursor-pointer flex justify-center items-center">
                            <div class="text-xs text-white font-bold"> FREE </div>
                        </div>
                    {/each}
                </div>
            {/each}
        {/each}
    </div>
</div>