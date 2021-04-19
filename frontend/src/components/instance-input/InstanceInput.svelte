<script lang=typescript>
    
    import type { InstanceData, Day } from "../../openapi";

    import { Service } from '../../openapi';

    import Modal from "../Modal.svelte";
    import Button from "../model-input/Button.svelte";
    import AssistantInput from "./AssistantInput.svelte";

    export let data: InstanceData

    function handleSubmit() {
        Service.postInstanceData(data)
    }

    function toggle_holiday(day: Day) {
        day.is_holiday = !day.is_holiday
        data.days = data.days
    }

    function removeWeek() {
        data.days = data.days.slice(0, data.days.length - 7)
    }

    function addWeek() {
        data.days.push(
            {
                id: data.days.length + 1,
                date: "todo",
                is_holiday: false
            },
            {
                id: data.days.length + 2,
                date: "todo",
                is_holiday: false
            },
            {
                id: data.days.length + 3,
                date: "todo",
                is_holiday: false
            },
            {
                id: data.days.length + 4,
                date: "todo",
                is_holiday: false
            },
            {
                id: data.days.length + 5,
                date: "todo",
                is_holiday: false
            },
            {
                id: data.days.length + 6,
                date: "todo",
                is_holiday: false
            },
            {
                id: data.days.length + 7,
                date: "todo",
                is_holiday: false
            },
        )
        console.log(data.days)
        data.days = data.days
    }

</script>

<main>
    <div class="flex flex-col space-y-2"> 
        <p class="font-semibold text-sm cursor-default"> Instance data </p>

        <p class="mt-4 font-semibold text-xs text-gray-500 cursor-default"> Number of weeks </p>
        <div class="flex flex-row justify-start space-x-4">
            <Button callback={() => removeWeek()}> - </Button>
            <p class="font-semibold text-sm text-black-500 cursor-default"> {data.days.length / 7} </p>
            <Button callback={() => addWeek()}> + </Button>
        </div>
        <Button callback={() => {handleSubmit()}}> Submit </Button>
        <Modal>
            <div slot="trigger" let:open> 
                <Button callback={open}> Edit holidays </Button>
            </div>
            <div class="my-2" slot="header">
                <h1 class="font-bold"> Edit holidays </h1>
            </div>
            <div class="my-2" slot="content">
                <div class="grid grid-cols-7 gap-4">
                    {#each data.days as day}
                        <div on:click={() => toggle_holiday(day)} class="cursor-pointer flex flex-row justify-items-center w-12 h-6 rounded {day.is_holiday ? "bg-yellow-200" : "bg-gray-100"}">
                            <p class="m-auto text-sm font-bold text-black"> {day.id} </p>
                        </div>
                    {/each}
                </div>
            </div>
            <div class="my-5 flex flex-row justify-end space-x-2" slot="footer" let:close>
                <Button primary={false} callback={close}> Close </Button>
                <Button callback={() => {handleSubmit(); close()}}> Submit </Button>
            </div>
        </Modal>
        <Modal>
            <div slot="trigger" let:open>
                <Button callback={open}> Edit assistants </Button>
            </div>
            <div class="my-2" slot="header">
                <h1 class="font-bold"> Edit assistants </h1>
            </div>
            <div class="my-2" slot="content">
                <AssistantInput bind:data={data}/>
            </div>
            <div class="my-5 flex flex-row justify-end space-x-2" slot="footer" let:close>
                <Button primary={false} callback={close}> Close </Button>
                <Button callback={() => {handleSubmit(); close()}}> Submit </Button>
            </div>
        </Modal>
    </div>
</main>