package packetPaddingContainer_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.torproject.org/tpo/anti-censorship/pluggable-transports/snowflake/v2/common/packetPaddingContainer"
)

func TestPacketPaddingContainer(t *testing.T) {
	Convey("Given a PacketPaddingContainer", t, func() {
		container := packetPaddingContainer.New()

		Convey("When packing data with padding", func() {
			data := []byte("testdata")
			padding := 4
			packedData := container.Pack(data, padding)

			Convey("The packed data should have the correct length", func() {
				expectedLength := len(data) + padding + 2
				So(len(packedData), ShouldEqual, expectedLength)
			})

			Convey("When unpacking the packed data", func() {
				unpackedData, unpackedPadding := container.Unpack(packedData)

				Convey("The unpacked data should match the original data", func() {
					So(string(unpackedData), ShouldEqual, string(data))
				})

				Convey("The unpacked padding should match the original padding", func() {
					So(unpackedPadding, ShouldEqual, padding)
				})
			})
		})

		Convey("When packing empty data with padding", func() {
			data := []byte("")
			padding := 4
			packedData := container.Pack(data, padding)

			Convey("The packed data should have the correct length", func() {
				expectedLength := len(data) + padding + 2
				So(len(packedData), ShouldEqual, expectedLength)
			})

			Convey("When unpacking the packed data", func() {
				unpackedData, unpackedPadding := container.Unpack(packedData)

				Convey("The unpacked data should match the original data", func() {
					So(string(unpackedData), ShouldEqual, string(data))
				})

				Convey("The unpacked padding should match the original padding", func() {
					So(unpackedPadding, ShouldEqual, padding)
				})
			})
		})

		Convey("When packing data with zero padding", func() {
			data := []byte("testdata")
			padding := 0
			packedData := container.Pack(data, padding)

			Convey("The packed data should have the correct length", func() {
				expectedLength := len(data) + padding + 2
				So(len(packedData), ShouldEqual, expectedLength)
			})

			Convey("When unpacking the packed data", func() {
				unpackedData, unpackedPadding := container.Unpack(packedData)

				Convey("The unpacked data should match the original data", func() {
					So(string(unpackedData), ShouldEqual, string(data))
				})

				Convey("The unpacked padding should match the original padding", func() {
					So(unpackedPadding, ShouldEqual, padding)
				})
			})
		})

		Convey("When padding data", func() {
			Convey("With a positive padding length", func() {
				padLength := 3
				padData := container.Pad(padLength)

				Convey("The padded data should have the correct length", func() {
					So(len(padData), ShouldEqual, padLength)
				})
			})

			Convey("With a zero padding length", func() {
				padLength := 0
				padData := container.Pad(padLength)

				Convey("The padded data should be empty", func() {
					So(len(padData), ShouldEqual, 0)
				})
			})

			Convey("With a negative padding length", func() {
				padLength := -1
				padData := container.Pad(padLength)

				Convey("The padded data should be nil", func() {
					So(padData, ShouldBeNil)
				})
			})
		})
	})
}
