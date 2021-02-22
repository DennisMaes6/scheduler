<script lang=typescript>

    import type {Assignment} from '../../openapi'
    import {ShiftType} from '../../openapi'

    export let assignment: Assignment

    function getColor(assignment: ShiftType): string {
        switch (assignment) {
            case ShiftType.JAEV:
                return "bg-yellow-400";
            case ShiftType.JANW:
                return "bg-green-400";
            case ShiftType.JAWH:
                return "bg-red-500";
            case ShiftType.SAEW:
                return "bg-green-500";
            case ShiftType.SAWH:
                return "bg-red-600";
            case ShiftType.CALL:
                return "bg-green-600";
            case ShiftType.TSPT:
                return "bg-red-700";
            case ShiftType.FREE:
                return "bg-white";
        }
    }

</script>

<main>
    {#if assignment.part_of_min_balance}
        <div class="flex w-12 h-6 rounded-lg justify-center items-center cursor-default">
            <div class="text-xs text-red-500 font-bold"> ! </div>
        </div>
    {:else}
        <div class="flex w-12 h-6 {getColor(assignment.shift_type)} rounded-lg justify-center items-center cursor-default">
            <div class="text-xs font-bold text-white"> {assignment.shift_type}</div>
        </div>
    {/if}
</main>