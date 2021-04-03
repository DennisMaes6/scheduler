
<script lang=typescript>

import type { AssistantInstance } from "../../openapi";
import { AssistantType } from "../../openapi";

import IconButton from "../model-input/IconButton.svelte";
import Button from "../model-input/Button.svelte";

import DayHeader from "../schedule-view/DayHeader.svelte";

export let assistants: AssistantInstance[]
export let nb_weeks: number;

let days: number[] = Array.from({length: nb_weeks*7}, (_, i) => i + 1)

const types: AssistantType[] = [
    AssistantType.JA,
    AssistantType.JA_F,
    AssistantType.SA,
    AssistantType.SA_F,
    AssistantType.SA_NEO,
    AssistantType.SA_F_NEO
]

function removeAssistant(assistant: AssistantInstance) {
        assistants = assistants.filter((a) => a.id != assistant.id)
        let index = 1
        assistants.forEach((a) => {
            a.id = index
            index++
        })
    }

    function addAssistant(type: AssistantType) {
        assistants.push(
            {
                id: Math.max(...assistants.map((a) => a.id)) + 1,
                type,
                free_days: []
            }
        )
        let index = 1
        assistants.forEach((a) => {
            a.id = index
            index++
        })
        assistants = assistants
    }
</script>


<div class="flex flex-col space-y-2">
    <div class="flex flex-row">
        <div class="w-20"/>
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
                <p class="w-16 text-black text-sm"> {assistant.type} </p>
                
            </div>
        {/each}
        <Button callback={() => addAssistant(type)}> New </Button>
    {/each}
</div>