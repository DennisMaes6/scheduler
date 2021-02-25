<script lang=typescript>
    
    import type { InstanceData, AssistantInstance} from "../../openapi";
    import { AssistantType } from "../../openapi";

    import { Service } from '../../openapi';

    import Button from "../model-input/Button.svelte";
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

</script>

<main>
    <form class="flex flex-col"> 
        <p class="font-semibold text-sm cursor-default"> Instance data </p>
        <p class="mt-2 font-semibold text-xs text-gray-500 cursor-default"> Number of weeks </p>
        <InputField bind:value={data.nb_weeks} />

        <p class="mt-2 font-semibold text-xs text-gray-500 cursor-default"> Number of assistants </p>
        <div class="mt-2 flex flex-col">
            {#each counts as countInstance}
                <div class="flex flex-row space-x-1 justify-between">
                    <p class="mx-2 w-12 text-black text-xs font-bold"> {countInstance.type} </p>
                    <div class=w-12>
                        <InputField bind:value={countInstance.count}/>
                    </div>
                </div>
            {/each}
        </div>

        <div class="my-5 mx-auto">
            <Button callback={handleSubmit}> Submit </Button>
        </div>
    </form>

</main>