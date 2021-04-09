<script lang="typescript">
    
    import type { Assistant } from "../../openapi";
    import { AssistantType } from "../../openapi";

    export let assistant: Assistant

    const types: AssistantType[] = Object.values(AssistantType)


</script>

<form>
    <label class="text-sm" for="name"> Name </label>
    <input class="text-sm pl-2" id="name" bind:value={assistant.name} placeholder="Jansen J"/>

    <label class="text-sm mt-3" for="type"> Type </label>
    <select class="text-sm pl-2" id="type" bind:value={assistant.type}>
        {#each types as type}
            <option class="text-sm pl-2" value={type}><p class="text-sm">{type}</p></option>
        {/each}
    </select>

</form>