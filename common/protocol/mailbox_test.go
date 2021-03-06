package protocol

import "testing"

func TestMailbox(t *testing.T) {
	mb := NewMailbox(0xFFF, 0, ID_MAX)
	if mb.ServiceId() != 0xFFF || mb.Flag() != 0 || mb.Id() != ID_MAX {
		t.Fatalf("mailbox error, want: 0xFFF, 0, ID_MAX, have:%d, %d, %d", mb.ServiceId(), mb.Flag(), mb.Id())
	}

	mb1 := NewMailbox(1, 1, 1000)
	if mb1.Flag() != 1 || mb1.Id() != 1000 {
		t.Fatalf("mailbox error, want: 1, 1000, have:%d, %d", mb1.Flag(), mb1.Id())
	}

	str := mb.String()
	mb3, err := NewMailboxFromStr(str)
	if err != nil {
		t.Fatal(err)
	}

	if mb3.ServiceId() != 0xFFF || mb3.Flag() != 0 || mb3.Id() != ID_MAX {
		t.Fatalf("mailbox error, want: 0xFFF, 0, ID_MAX, have:%d, %d, %d", mb3.ServiceId(), mb3.Flag(), mb3.Id())
	}

	mb4 := NewMailboxFromUid(mb1.Uid())
	if mb4.Flag() != 1 || mb4.Id() != 1000 {
		t.Fatalf("mailbox error, want: 1, 1000, have:%d, %d", mb4.Flag(), mb4.Id())
	}

}
