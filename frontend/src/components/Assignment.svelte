<script lang=typescript>

    import {ShiftType} from '../openapi'

    export let assignment: ShiftType

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
    <div class="flex w-16 h-8 {getColor(assignment)} rounded-lg justify-center items-center">
        <p class="text-sm font-bold text-white"> {assignment}</p>
    </div>
</main>