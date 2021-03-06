package cmd

import (
	"errors"
	"github.com/cheynewallace/tabby"
	"github.com/leanovate/mite-go/domain"
	"github.com/spf13/cobra"
)

var (
	trackerTimeEntryId string
)

func init() {
	trackerCommand.AddCommand(trackerStatusCommand)
	trackerStartCommand.Flags().StringVarP(&trackerTimeEntryId, "id", "i", "", "the time entry id to (re)start a tracker for (default: latest time entry for today)")
	trackerCommand.AddCommand(trackerStartCommand)
	trackerStopCommand.Flags().StringVarP(&trackerTimeEntryId, "id", "i", "", "the time entry id to stop a tracker for (default: latest time entry for today)")
	trackerCommand.AddCommand(trackerStopCommand)
	rootCmd.AddCommand(trackerCommand)
}

var trackerCommand = &cobra.Command{
	Use:   "tracker",
	Short: "starts, stops and shows the status of the tracker",
	RunE:  trackerStatusCommand.RunE,
}

var trackerStatusCommand = &cobra.Command{
	Use:   "status",
	Short: "shows the status of the time tracker",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		tracking, err := application.MiteApi.Tracker()
		if err != nil {
			return err
		}
		if tracking == nil {
			return nil
		}

		t := tabby.New()
		t.AddHeader("id", "time", "state", "since")
		t.AddLine(tracking.Id, tracking.Minutes, "tracking", tracking.Since.In(application.Conf.GetDisplayLocation()))
		t.Print()

		return nil
	},
}

var trackerStartCommand = &cobra.Command{
	Use:   "start",
	Short: "starts the time tracker for a time entry",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var entryId domain.TimeEntryId
		if trackerTimeEntryId == "" {
			entryId, err = fetchLatestTimeEntryForToday()
		} else {
			entryId, err = domain.ParseTimeEntryId(trackerTimeEntryId)
		}
		if err != nil {
			return err
		}

		tracking, stopped, err := application.MiteApi.StartTracker(entryId)
		if err != nil {
			return err
		}

		t := tabby.New()
		t.AddHeader("id", "time", "state", "since")
		t.AddLine(tracking.Id, tracking.Minutes, "tracking", tracking.Since.In(application.Conf.GetDisplayLocation()))
		if stopped != nil {
			t.AddLine(stopped.Id, stopped.Minutes, "stopped")
		}
		t.Print()

		return nil
	},
}

var trackerStopCommand = &cobra.Command{
	Use:   "stop",
	Short: "stops the time tracker for a time entry",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var entryId domain.TimeEntryId
		if trackerTimeEntryId == "" {
			entryId, err = fetchLatestTimeEntryForToday()
		} else {
			entryId, err = domain.ParseTimeEntryId(trackerTimeEntryId)
		}
		if err != nil {
			return err
		}

		stopped, err := application.MiteApi.StopTracker(entryId)
		if err != nil {
			return err
		}

		t := tabby.New()
		t.AddHeader("id", "time", "state")
		t.AddLine(stopped.Id, stopped.Minutes, "stopped")
		t.Print()

		return nil
	},
}

func fetchLatestTimeEntryForToday() (domain.TimeEntryId, error) {
	today := domain.Today()

	entries, err := application.MiteApi.TimeEntries(&domain.TimeEntryQuery{
		To:        &today,
		From:      &today,
		Direction: "desc",
	})
	if err != nil {
		return 0, err
	}

	if len(entries) == 0 {
		return 0, errors.New("no time entries for today found")
	}

	return entries[0].Id, nil
}
