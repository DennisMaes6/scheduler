<script lang=typescript>
    
    import type { InstanceData, AssistantInstance} from "../../openapi";
    import { AssistantType } from "../../openapi";

    import { Service } from '../../openapi';

    import Button from "../model-input/Button.svelte";
    import IconButton from "../model-input/IconButton.svelte";
    import InputField from "../model-input/InputField.svelte";

    export let data: InstanceData

    function count(type: AssistantType): number {
        return data.assistants.filter((ai) => ai.type === type).length
    }

    let counts = [
        {
            type: AssistantType.JA,
            count: count(AssistantType.JA),
        },
        {
            type: AssistantType.JA_F,
            count: count(AssistantType.JA_F),
        },
        {
            type: AssistantType.SA,
            count: count(AssistantType.SA),
        },
        {
            type: AssistantType.SA_F,
            count: count(AssistantType.SA_F),
        },
        {
            type: AssistantType.SA_NEO,
            count: count(AssistantType.SA_NEO),
        },
        {
            type: AssistantType.SA_F_NEO,
            count: count(AssistantType.SA_F_NEO),
        },
    ]

    function handleSubmit() {
        let index = 0
        let newData: AssistantInstance[] = [] 
        counts.forEach((countInstance) => {
            for (let j = 0; j < countInstance.count; j++) {
                index++
                newData.push({
                    id: index,
                    type: countInstance.type
                })
            }
        })
        data.assistants = newData
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

        <p class="mt-1 font-semibold text-xs text-gray-500 cursor-default"> Number of assistants </p>
        <div class="mt-1 flex flex-col">
            {#each counts as countInstance}
                <div class="flex flex-row space-x-1 justify-between">
                    <p class="mx-2 w-12 text-black text-xs font-bold"> {countInstance.type} </p>
                    <div class=w-12>
                        <InputField bind:value={countInstance.count}/>
                    </div>
                </div>
            {/each}
        </div>

        <div class="my-1 mx-auto">
            <Button callback={handleSubmit}> Submit </Button>
        </div>
    </form>

</main>