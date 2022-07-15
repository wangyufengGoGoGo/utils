package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestSecondDuration(t *testing.T) {
	duration := SecondDuration(60)
	second := 60 * time.Second
	if duration != second {
		t.Fatalf("duration error, %s not equal %s", duration, second)
	}
	fmt.Println(duration)
}

func TestMinuteDuration(t *testing.T) {
	duration := MinuteDuration(60)
	minute := 60 * time.Minute
	if duration != minute {
		t.Fatalf("duration error, %s not equal %s", duration, minute)
	}
	fmt.Println(duration)
}

func TestHourDuration(t *testing.T) {
	duration := HourDuration(60)
	hour := 60 * time.Hour
	if duration != hour {
		t.Fatalf("duration error, %s not equal %s", duration, hour)
	}
	fmt.Println(duration)
}

func TestAfter(t *testing.T) {
	t.Run("当前时间之前", func(t *testing.T) {
		start, err := time.Parse(TimeFormat, "2022-07-08 15:04:05")
		if err != nil {
			t.Fatalf("parse time error: {%s}", err)
		}

		if After(start) {
			t.Log("start should more than now")
		}
		t.Log("start should less than now")
	})
	t.Run("当前时间之前(1657853331)", func(t *testing.T) {
		start, err := time.Parse(TimeFormat, "2022-07-15 10:48:51")
		if err != nil {
			t.Fatalf("parse time error: {%s}", err)
		}
		if After(start) {
			t.Log("start should more than now")
		}
		t.Log("start should less than now")
	})
	t.Run("当前时间之后", func(t *testing.T) {
		if After(time.Now()) {
			t.Log("start should more than now")
		}
		t.Log("start should less than now")
	})
}

func TestSecondAfter(t *testing.T) {
	t.Run("当前时间之前", func(t *testing.T) {
		start, err := time.Parse(TimeFormat, "2022-07-08 15:04:05")
		if err != nil {
			t.Fatalf("parse time error: {%s}", err)
		}

		if SecondAfter(start.Unix(), 60) {
			t.Fatalf("start less than now")
		}
	})
	t.Run("当前时间之后", func(t *testing.T) {
		if !SecondAfter(time.Now().Unix(), 60) {
			t.Fatalf("start less than now")
		}
	})
}

func TestMinuteAfter(t *testing.T) {
	t.Run("当前时间之前", func(t *testing.T) {
		start, err := time.Parse(TimeFormat, "2022-07-08 15:04:05")
		if err != nil {
			t.Fatalf("parse time error: {%s}", err)
		}

		if MinuteAfter(start.Unix(), 60) {
			t.Fatalf("start less than now")
		}
	})
	t.Run("当前时间之后", func(t *testing.T) {
		if !MinuteAfter(time.Now().Unix(), 60) {
			t.Fatalf("start less than now")
		}
	})
}

func TestHourAfter(t *testing.T) {
	t.Run("当前时间之前", func(t *testing.T) {
		start, err := time.Parse(TimeFormat, "2022-07-08 15:04:05")
		if err != nil {
			t.Fatalf("parse time error: {%s}", err)
		}

		if HourAfter(start.Unix(), 60) {
			t.Fatalf("start should less than now")
		}
	})
	t.Run("当前时间之后", func(t *testing.T) {
		if !HourAfter(time.Now().Unix(), 60) {
			t.Fatalf("start should more than now")
		}
	})
}

func TestAfterByUnix(t *testing.T) {
	t.Run("当前时间之前", func(t *testing.T) {
		start, err := time.Parse(TimeFormat, "2022-07-08 15:04:05")
		if err != nil {
			t.Fatalf("parse time error: {%s}", err)
		}

		if AfterByUnix(start.Unix(), 60*time.Second) {
			t.Fatalf("start less than now")
		}
	})
	t.Run("当前时间之前(1657853331)", func(t *testing.T) {
		//2022-07-15 10:48:51
		unix := int64(1657853331)
		if AfterByUnix(unix, 60*time.Second) {
			t.Fatalf("start less than now")
		}
	})
	t.Run("当前时间之后", func(t *testing.T) {
		if !AfterByUnix(time.Now().Unix(), 60*time.Second) {
			t.Fatalf("start less than now")
		}
	})
}
